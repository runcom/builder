#!/usr/bin/env bats

load helpers

@test "blobcache-pull" {
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	# Pull an image using a fresh directory for the blob cache.
	run buildah pull --blob-cache=${blobcachedir} docker.io/kubernetes/pause
	echo "$output"
	[ "$status" -eq 0 ]
	# Check that we dropped some files in there.
	run find ${blobcachedir} -type f
	echo "$output"
	[ "$status" -eq 0 ]
	[ "${#lines[@]}" -gt 0 ]
}

@test "blobcache-from" {
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	# Pull an image using a fresh directory for the blob cache.
	run buildah from --blob-cache=${blobcachedir} docker.io/kubernetes/pause
	echo "$output"
	[ "$status" -eq 0 ]
	# Check that we dropped some files in there.
	run find ${blobcachedir} -type f
	echo "$output"
	[ "$status" -eq 0 ]
	[ "${#lines[*]}" -gt 0 ]
}

@test "blobcache-commit" {
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	# Pull an image using a fresh directory for the blob cache.
	run buildah --debug=false from --quiet --blob-cache=${blobcachedir} docker.io/kubernetes/pause
	echo "$output"
	ctr="$output"
	[ "$status" -eq 0 ]
	# Commit the image using the blob cache.
	destdir=${TESTDIR}/dest
	mkdir -p ${destdir}
	echo buildah commit --blob-cache=${blobcachedir} ${ctr} dir:${destdir}
	run buildah commit --blob-cache=${blobcachedir} ${ctr} dir:${destdir}
	echo "$output"
	[ "$status" -eq 0 ]
	# Look for layer blobs in the destination that match the ones in the cache.
	matched=0
	unmatched=0
	for blob in ${blobcachedir}/* ; do
		match=false
		for content in ${destdir}/* ; do
			if cmp -s ${content} ${blob} ; then
				match=true
				break
			fi
		done
		if ${match} ; then
			matched=$(( ${matched} + 1 ))
		else
			unmatched=$(( ${unmatched} + 1 ))
		fi
	done
	# We shouldn't have any matches, since we had to recompress everything, and the configuration
	# has more history in it.
	[ ${matched} -gt 0 ]
	[ ${unmatched} -le 1 ]
}

@test "blobcache-push" {
	target=targetimage
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	# Pull an image using a fresh directory for the blob cache.
	run buildah --debug=false from --quiet --blob-cache=${blobcachedir} docker.io/kubernetes/pause
	echo "$output"
	ctr="$output"
	[ "$status" -eq 0 ]
	# Commit the image using the blob cache.
	destdir=${TESTDIR}/dest
	mkdir -p ${destdir}
	run buildah commit --blob-cache=${blobcachedir} ${ctr} ${target}
	echo "$output"
	[ "$status" -eq 0 ]
	# Try to push the image without the blob cache.
	doomeddir=${TESTDIR}/doomed
	mkdir -p ${doomeddir}
	run buildah push ${target} dir:${doomeddir}
	echo "$output"
	[ "$status" -eq 0 ]
	# Look for layer blobs in the doomed copy that match the ones in the cache.
	matched=0
	unmatched=0
	for content in ${doomeddir}/* ; do
		match=false
		for blob in ${blobcachedir}/* ; do
			if cmp -s ${content} ${blob} ; then
				match=true
				break
			fi
		done
		if ${match} ; then
			matched=$(( ${matched} + 1 ))
			echo ${content} was cached
		else
			unmatched=$(( ${unmatched} + 1 ))
			echo ${content} was not cached
		fi
	done
	# We should have no matches.
	[ ${matched} -eq 0 ]
	[ ${unmatched} -gt 1 ]
	# Now try to push the image with the blob cache.
	destdir=${TESTDIR}/dest
	mkdir -p ${destdir}
	run buildah push --blob-cache=${blobcachedir} ${target} dir:${destdir}
	echo "$output"
	[ "$status" -eq 0 ]
	# Look for layer blobs in the destination that match the ones in the cache.
	matched=0
	unmatched=0
	for content in ${destdir}/* ; do
		match=false
		for blob in ${blobcachedir}/* ; do
			if cmp -s ${content} ${blob} ; then
				match=true
				break
			fi
		done
		if ${match} ; then
			matched=$(( ${matched} + 1 ))
			echo ${content} was cached
		else
			unmatched=$(( ${unmatched} + 1 ))
			echo ${content} was not cached
		fi
	done
	# We should have at most three mismatches and at least one match.
	[ ${matched} -gt 0 ]
	[ ${unmatched} -le 3 ] # the expected three are the "version", manifest, config
}

@test "blobcache-build-using-dockerfile-local" {
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	target=new-image
	# Build an image while pulling the base image.
	run buildah build-using-dockerfile -t ${target} --pull-always --blob-cache=${blobcachedir} ${TESTSDIR}/bud/add-file
	echo "$output"
	[ "$status" -eq 0 ]
	# Now try to push the image with the blob cache.
	destdir=${TESTDIR}/dest
	mkdir -p ${destdir}
	run buildah push --blob-cache=${blobcachedir} ${target} dir:${destdir}
	echo "$output"
	[ "$status" -eq 0 ]
	# Look for layer blobs in the destination that match the ones in the cache.
	matched=0
	unmatched=0
	for content in ${destdir}/* ; do
		match=false
		for blob in ${blobcachedir}/* ; do
			if cmp -s ${content} ${blob} ; then
				match=true
				break
			fi
		done
		if ${match} ; then
			matched=$(( ${matched} + 1 ))
			echo ${content} was cached
		else
			unmatched=$(( ${unmatched} + 1 ))
			echo ${content} was not cached
		fi
	done
	# We should have at most three mismatches and at least one match.
	[ ${matched} -gt 0 ]
	[ ${unmatched} -le 3 ] # the expected three are the "version", manifest, config
}

@test "blobcache-build-using-dockerfile-push" {
	blobcachedir=${TESTDIR}/cache
	mkdir -p ${blobcachedir}
	target=new-image
	destdir=${TESTDIR}/dest
	mkdir -p ${destdir}
	# Build an image while pulling the base image, implicitly pushing while writing.
	run buildah build-using-dockerfile -t dir:${destdir} --pull-always --blob-cache=${blobcachedir} ${TESTSDIR}/bud/add-file
	echo "$output"
	[ "$status" -eq 0 ]
	# Look for layer blobs in the destination that match the ones in the cache.
	matched=0
	unmatched=0
	for content in ${destdir}/* ; do
		match=false
		for blob in ${blobcachedir}/* ; do
			if cmp -s ${content} ${blob} ; then
				match=true
				break
			fi
		done
		if ${match} ; then
			matched=$(( ${matched} + 1 ))
			echo ${content} was cached
		else
			unmatched=$(( ${unmatched} + 1 ))
			echo ${content} was not cached
		fi
	done
	# We should have at most three mismatches and at least one match.
	[ ${matched} -gt 0 ]
	[ ${unmatched} -le 3 ] # the expected three are the "version", manifest, config
}
