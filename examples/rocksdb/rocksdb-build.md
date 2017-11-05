# Facebook rocksdb

## Second Build
Pull
```
[vagrant@localhost rocksdb]$ git pull --append
remote: Counting objects: 14076, done.
remote: Compressing objects: 100% (4412/4412), done.
remote: Total 14076 (delta 10827), reused 12130 (delta 8985), pack-reused 0
接收对象中: 100% (14076/14076), 7.01 MiB | 319.00 KiB/s, 完成.
处理 delta 中: 100% (10827/10827), 完成 651 个本地对象.
来自 https://github.com/facebook/rocksdb
   88a2776..7318952  master     -> origin/master
 * [新tag]           rocksdb-5.8 -> rocksdb-5.8
 * [新tag]           v5.7.1     -> v5.7.1
 * [新tag]           v5.8       -> v5.8
更新 88a2776..7318952
正在检出文件: 100% (1285/1285), 完成.
Fast-forward
### snippets ###
```

Revision
```
[vagrant@localhost rocksdb]$ git log -1
commit 731895214bb729f4b748c8051d2d4c81927ff47f
Author: Andrew Kryczka <andrewkr@fb.com>
Date:   Mon Oct 16 18:29:50 2017 -0700

    db_bench randomtransaction print throughput
    
    Summary:
    print throughput in MB/s upon finishing randomtransaction benchmark
    Closes https://github.com/facebook/rocksdb/pull/3016
    
    Differential Revision: D6070426
    
    Pulled By: ajkr
    
    fbshipit-source-id: 69df43beed4c374a36d826e761ca3a83e1fdcbf5
```

### Build static_lib
Make
```
[vagrant@localhost rocksdb]$ make static_lib
  GEN      util/build_version.cc
  GEN      util/build_version.cc
  CC       cache/clock_cache.o
  CC       cache/lru_cache.o
  CC       cache/sharded_cache.o
  CC       db/builder.o
  CC       db/c.o
  CC       db/column_family.o
  CC       db/compacted_db_impl.o
  CC       db/compaction.o
  CC       db/compaction_iterator.o
  CC       db/compaction_job.o
  CC       db/compaction_picker.o
  CC       db/compaction_picker_universal.o
  CC       db/convenience.o
  CC       db/db_filesnapshot.o
  CC       db/db_impl.o
  CC       db/db_impl_write.o
  CC       db/db_impl_compaction_flush.o
  CC       db/db_impl_files.o
  CC       db/db_impl_open.o
  CC       db/db_impl_debug.o
  CC       db/db_impl_experimental.o
  CC       db/db_impl_readonly.o
  CC       db/db_info_dumper.o
  CC       db/db_iter.o
  CC       db/dbformat.o
  CC       db/event_helpers.o
  CC       db/experimental.o
  CC       db/external_sst_file_ingestion_job.o
  CC       db/file_indexer.o
  CC       db/flush_job.o
  CC       db/flush_scheduler.o
  CC       db/forward_iterator.o
  CC       db/internal_stats.o
  CC       db/log_reader.o
  CC       db/log_writer.o
  CC       db/malloc_stats.o
  CC       db/managed_iterator.o
  CC       db/memtable.o
  CC       db/memtable_list.o
  CC       db/merge_helper.o
  CC       db/merge_operator.o
  CC       db/range_del_aggregator.o
  CC       db/repair.o
  CC       db/snapshot_impl.o
  CC       db/table_cache.o
  CC       db/table_properties_collector.o
  CC       db/transaction_log_impl.o
  CC       db/version_builder.o
  CC       db/version_edit.o
  CC       db/version_set.o
  CC       db/wal_manager.o
  CC       db/write_batch.o
  CC       db/write_batch_base.o
  CC       db/write_controller.o
  CC       db/write_thread.o
  CC       env/env.o
  CC       env/env_chroot.o
  CC       env/env_encryption.o
  CC       env/env_hdfs.o
  CC       env/env_posix.o
  CC       env/io_posix.o
  CC       env/mock_env.o
  CC       memtable/alloc_tracker.o
  CC       memtable/hash_cuckoo_rep.o
  CC       memtable/hash_linklist_rep.o
  CC       memtable/hash_skiplist_rep.o
  CC       memtable/skiplistrep.o
  CC       memtable/vectorrep.o
  CC       memtable/write_buffer_manager.o
  CC       monitoring/histogram.o
  CC       monitoring/histogram_windowing.o
  CC       monitoring/instrumented_mutex.o
  CC       monitoring/iostats_context.o
  CC       monitoring/perf_context.o
  CC       monitoring/perf_level.o
  CC       monitoring/statistics.o
  CC       monitoring/thread_status_impl.o
  CC       monitoring/thread_status_updater.o
  CC       monitoring/thread_status_updater_debug.o
  CC       monitoring/thread_status_util.o
  CC       monitoring/thread_status_util_debug.o
  CC       options/cf_options.o
  CC       options/db_options.o
  CC       options/options.o
  CC       options/options_helper.o
  CC       options/options_parser.o
  CC       options/options_sanity_check.o
  CC       port/port_posix.o
  CC       port/stack_trace.o
  CC       table/adaptive_table_factory.o
  CC       table/block.o
  CC       table/block_based_filter_block.o
  CC       table/block_based_table_builder.o
  CC       table/block_based_table_factory.o
  CC       table/block_based_table_reader.o
  CC       table/block_builder.o
  CC       table/block_prefix_index.o
  CC       table/bloom_block.o
  CC       table/cuckoo_table_builder.o
  CC       table/cuckoo_table_factory.o
  CC       table/cuckoo_table_reader.o
  CC       table/flush_block_policy.o
  CC       table/format.o
  CC       table/full_filter_block.o
  CC       table/get_context.o
  CC       table/index_builder.o
  CC       table/iterator.o
  CC       table/merging_iterator.o
  CC       table/meta_blocks.o
  CC       table/partitioned_filter_block.o
  CC       table/persistent_cache_helper.o
  CC       table/plain_table_builder.o
  CC       table/plain_table_factory.o
  CC       table/plain_table_index.o
  CC       table/plain_table_key_coding.o
  CC       table/plain_table_reader.o
  CC       table/sst_file_writer.o
  CC       table/table_properties.o
  CC       table/two_level_iterator.o
  CC       tools/dump/db_dump_tool.o
  CC       util/arena.o
  CC       util/auto_roll_logger.o
  CC       util/bloom.o
  CC       util/build_version.o
  CC       util/coding.o
  CC       util/compaction_job_stats_impl.o
  CC       util/comparator.o
  CC       util/concurrent_arena.o
  CC       util/crc32c.o
  CC       util/delete_scheduler.o
  CC       util/dynamic_bloom.o
  CC       util/event_logger.o
  CC       util/file_reader_writer.o
  CC       util/file_util.o
  CC       util/filename.o
  CC       util/filter_policy.o
  CC       util/hash.o
  CC       util/log_buffer.o
  CC       util/murmurhash.o
  CC       util/random.o
  CC       util/rate_limiter.o
  CC       util/slice.o
  CC       util/sst_file_manager_impl.o
  CC       util/status.o
  CC       util/status_message.o
  CC       util/string_util.o
  CC       util/sync_point.o
  CC       util/thread_local.o
  CC       util/threadpool_imp.o
  CC       util/transaction_test_util.o
  CC       util/xxhash.o
  CC       utilities/backupable/backupable_db.o
  CC       utilities/blob_db/blob_db.o
  CC       utilities/blob_db/blob_db_impl.o
  CC       utilities/blob_db/blob_file.o
  CC       utilities/blob_db/blob_log_reader.o
  CC       utilities/blob_db/blob_log_writer.o
  CC       utilities/blob_db/blob_log_format.o
  CC       utilities/blob_db/ttl_extractor.o
  CC       utilities/cassandra/cassandra_compaction_filter.o
  CC       utilities/cassandra/format.o
  CC       utilities/cassandra/merge_operator.o
  CC       utilities/checkpoint/checkpoint_impl.o
  CC       utilities/compaction_filters/remove_emptyvalue_compactionfilter.o
  CC       utilities/convenience/info_log_finder.o
  CC       utilities/date_tiered/date_tiered_db_impl.o
  CC       utilities/debug.o
  CC       utilities/document/document_db.o
  CC       utilities/document/json_document.o
  CC       utilities/document/json_document_builder.o
  CC       utilities/env_mirror.o
  CC       utilities/env_timed.o
  CC       utilities/geodb/geodb_impl.o
  CC       utilities/leveldb_options/leveldb_options.o
  CC       utilities/lua/rocks_lua_compaction_filter.o
  CC       utilities/memory/memory_util.o
  CC       utilities/merge_operators/max.o
  CC       utilities/merge_operators/put.o
  CC       utilities/merge_operators/string_append/stringappend.o
  CC       utilities/merge_operators/string_append/stringappend2.o
  CC       utilities/merge_operators/uint64add.o
  CC       utilities/option_change_migration/option_change_migration.o
  CC       utilities/options/options_util.o
  CC       utilities/persistent_cache/block_cache_tier.o
  CC       utilities/persistent_cache/block_cache_tier_file.o
  CC       utilities/persistent_cache/block_cache_tier_metadata.o
  CC       utilities/persistent_cache/persistent_cache_tier.o
  CC       utilities/persistent_cache/volatile_tier_impl.o
  CC       utilities/redis/redis_lists.o
  CC       utilities/simulator_cache/sim_cache.o
  CC       utilities/spatialdb/spatial_db.o
  CC       utilities/table_properties_collectors/compact_on_deletion_collector.o
  CC       utilities/transactions/optimistic_transaction_db_impl.o
  CC       utilities/transactions/optimistic_transaction.o
  CC       utilities/transactions/pessimistic_transaction.o
  CC       utilities/transactions/pessimistic_transaction_db.o
  CC       utilities/transactions/snapshot_checker.o
  CC       utilities/transactions/transaction_base.o
  CC       utilities/transactions/transaction_db_mutex_impl.o
  CC       utilities/transactions/transaction_lock_mgr.o
  CC       utilities/transactions/transaction_util.o
  CC       utilities/transactions/write_prepared_txn.o
  CC       utilities/ttl/db_ttl_impl.o
  CC       utilities/write_batch_with_index/write_batch_with_index.o
  CC       utilities/write_batch_with_index/write_batch_with_index_internal.o
  CC       tools/ldb_cmd.o
  CC       tools/ldb_tool.o
  CC       tools/sst_dump_tool.o
  CC       utilities/blob_db/blob_dump_tool.o
  AR       librocksdb.a
ar: 正在创建 librocksdb.a
```

Install
```
[vagrant@localhost rocksdb]$ sudo make install static_lib
Makefile:127: Warning: Compiling in debug mode. Don't use the resulting binary in production
make: Warning: File 'make_config.mk' has modification time 676 s in the future
  GEN      util/build_version.cc
install -d /usr/local/lib
for header_dir in `find "include/rocksdb" -type d`; do \
	install -d /usr/local/$header_dir; \
done
for header in `find "include/rocksdb" -type f -name *.h`; do \
	install -C -m 644 $header /usr/local/$header; \
done
  AR       librocksdb_debug.a
ar: 正在创建 librocksdb_debug.a
install -C -m 755 librocksdb_debug.a /usr/local/lib
[ -e librocksdb_debug.so.5.8.0 ] && make install-shared || :
make: Nothing to be done for 'static_lib'.
make: 警告：检测到时钟错误。您的创建可能是不完整的。
```

Run `ls`
```
[vagrant@localhost rocksdb]$ ls /usr/local/lib/librocksdb*
/usr/local/lib/librocksdb.a        /usr/local/lib/librocksdb.so    /usr/local/lib/librocksdb.so.4.9
/usr/local/lib/librocksdb_debug.a  /usr/local/lib/librocksdb.so.4  /usr/local/lib/librocksdb.so.4.9.0
```

### Build shared_lib
Make
```
[vagrant@localhost rocksdb]$ make shared_lib
make: Warning: File 'make_config.mk' has modification time 676 s in the future
  GEN      util/build_version.cc
  CC       shared-objects/cache/clock_cache.o
  CC       shared-objects/cache/lru_cache.o
  CC       shared-objects/cache/sharded_cache.o
  CC       shared-objects/db/builder.o
  CC       shared-objects/db/c.o
  CC       shared-objects/db/column_family.o
  CC       shared-objects/db/compacted_db_impl.o
  CC       shared-objects/db/compaction.o
  CC       shared-objects/db/compaction_iterator.o
  CC       shared-objects/db/compaction_job.o
  CC       shared-objects/db/compaction_picker.o
  CC       shared-objects/db/compaction_picker_universal.o
  CC       shared-objects/db/convenience.o
  CC       shared-objects/db/db_filesnapshot.o
  CC       shared-objects/db/db_impl.o
  CC       shared-objects/db/db_impl_write.o
  CC       shared-objects/db/db_impl_compaction_flush.o
  CC       shared-objects/db/db_impl_files.o
  CC       shared-objects/db/db_impl_open.o
  CC       shared-objects/db/db_impl_debug.o
  CC       shared-objects/db/db_impl_experimental.o
  CC       shared-objects/db/db_impl_readonly.o
  CC       shared-objects/db/db_info_dumper.o
  CC       shared-objects/db/db_iter.o
  CC       shared-objects/db/dbformat.o
  CC       shared-objects/db/event_helpers.o
  CC       shared-objects/db/experimental.o
  CC       shared-objects/db/external_sst_file_ingestion_job.o
  CC       shared-objects/db/file_indexer.o
  CC       shared-objects/db/flush_job.o
  CC       shared-objects/db/flush_scheduler.o
  CC       shared-objects/db/forward_iterator.o
  CC       shared-objects/db/internal_stats.o
  CC       shared-objects/db/log_reader.o
  CC       shared-objects/db/log_writer.o
  CC       shared-objects/db/malloc_stats.o
  CC       shared-objects/db/managed_iterator.o
  CC       shared-objects/db/memtable.o
  CC       shared-objects/db/memtable_list.o
  CC       shared-objects/db/merge_helper.o
  CC       shared-objects/db/merge_operator.o
  CC       shared-objects/db/range_del_aggregator.o
  CC       shared-objects/db/repair.o
  CC       shared-objects/db/snapshot_impl.o
  CC       shared-objects/db/table_cache.o
  CC       shared-objects/db/table_properties_collector.o
  CC       shared-objects/db/transaction_log_impl.o
  CC       shared-objects/db/version_builder.o
  CC       shared-objects/db/version_edit.o
  CC       shared-objects/db/version_set.o
  CC       shared-objects/db/wal_manager.o
  CC       shared-objects/db/write_batch.o
  CC       shared-objects/db/write_batch_base.o
  CC       shared-objects/db/write_controller.o
  CC       shared-objects/db/write_thread.o
  CC       shared-objects/env/env.o
  CC       shared-objects/env/env_chroot.o
  CC       shared-objects/env/env_encryption.o
  CC       shared-objects/env/env_hdfs.o
  CC       shared-objects/env/env_posix.o
  CC       shared-objects/env/io_posix.o
  CC       shared-objects/env/mock_env.o
  CC       shared-objects/memtable/alloc_tracker.o
  CC       shared-objects/memtable/hash_cuckoo_rep.o
  CC       shared-objects/memtable/hash_linklist_rep.o
  CC       shared-objects/memtable/hash_skiplist_rep.o
  CC       shared-objects/memtable/skiplistrep.o
  CC       shared-objects/memtable/vectorrep.o
  CC       shared-objects/memtable/write_buffer_manager.o
  CC       shared-objects/monitoring/histogram.o
  CC       shared-objects/monitoring/histogram_windowing.o
  CC       shared-objects/monitoring/instrumented_mutex.o
  CC       shared-objects/monitoring/iostats_context.o
  CC       shared-objects/monitoring/perf_context.o
  CC       shared-objects/monitoring/perf_level.o
  CC       shared-objects/monitoring/statistics.o
  CC       shared-objects/monitoring/thread_status_impl.o
  CC       shared-objects/monitoring/thread_status_updater.o
  CC       shared-objects/monitoring/thread_status_updater_debug.o
  CC       shared-objects/monitoring/thread_status_util.o
  CC       shared-objects/monitoring/thread_status_util_debug.o
  CC       shared-objects/options/cf_options.o
  CC       shared-objects/options/db_options.o
  CC       shared-objects/options/options.o
  CC       shared-objects/options/options_helper.o
  CC       shared-objects/options/options_parser.o
  CC       shared-objects/options/options_sanity_check.o
  CC       shared-objects/port/port_posix.o
  CC       shared-objects/port/stack_trace.o
  CC       shared-objects/table/adaptive_table_factory.o
  CC       shared-objects/table/block.o
  CC       shared-objects/table/block_based_filter_block.o
  CC       shared-objects/table/block_based_table_builder.o
  CC       shared-objects/table/block_based_table_factory.o
  CC       shared-objects/table/block_based_table_reader.o
  CC       shared-objects/table/block_builder.o
  CC       shared-objects/table/block_prefix_index.o
  CC       shared-objects/table/bloom_block.o
  CC       shared-objects/table/cuckoo_table_builder.o
  CC       shared-objects/table/cuckoo_table_factory.o
  CC       shared-objects/table/cuckoo_table_reader.o
  CC       shared-objects/table/flush_block_policy.o
  CC       shared-objects/table/format.o
  CC       shared-objects/table/full_filter_block.o
  CC       shared-objects/table/get_context.o
  CC       shared-objects/table/index_builder.o
  CC       shared-objects/table/iterator.o
  CC       shared-objects/table/merging_iterator.o
  CC       shared-objects/table/meta_blocks.o
  CC       shared-objects/table/partitioned_filter_block.o
  CC       shared-objects/table/persistent_cache_helper.o
  CC       shared-objects/table/plain_table_builder.o
  CC       shared-objects/table/plain_table_factory.o
  CC       shared-objects/table/plain_table_index.o
  CC       shared-objects/table/plain_table_key_coding.o
  CC       shared-objects/table/plain_table_reader.o
  CC       shared-objects/table/sst_file_writer.o
  CC       shared-objects/table/table_properties.o
  CC       shared-objects/table/two_level_iterator.o
  CC       shared-objects/tools/dump/db_dump_tool.o
  CC       shared-objects/util/arena.o
  CC       shared-objects/util/auto_roll_logger.o
  CC       shared-objects/util/bloom.o
  CC       shared-objects/util/build_version.o
  CC       shared-objects/util/coding.o
  CC       shared-objects/util/compaction_job_stats_impl.o
  CC       shared-objects/util/comparator.o
  CC       shared-objects/util/concurrent_arena.o
  CC       shared-objects/util/crc32c.o
  CC       shared-objects/util/delete_scheduler.o
  CC       shared-objects/util/dynamic_bloom.o
  CC       shared-objects/util/event_logger.o
  CC       shared-objects/util/file_reader_writer.o
  CC       shared-objects/util/file_util.o
  CC       shared-objects/util/filename.o
  CC       shared-objects/util/filter_policy.o
  CC       shared-objects/util/hash.o
  CC       shared-objects/util/log_buffer.o
  CC       shared-objects/util/murmurhash.o
  CC       shared-objects/util/random.o
  CC       shared-objects/util/rate_limiter.o
  CC       shared-objects/util/slice.o
  CC       shared-objects/util/sst_file_manager_impl.o
  CC       shared-objects/util/status.o
  CC       shared-objects/util/status_message.o
  CC       shared-objects/util/string_util.o
  CC       shared-objects/util/sync_point.o
  CC       shared-objects/util/thread_local.o
  CC       shared-objects/util/threadpool_imp.o
  CC       shared-objects/util/transaction_test_util.o
  CC       shared-objects/util/xxhash.o
  CC       shared-objects/utilities/backupable/backupable_db.o
  CC       shared-objects/utilities/blob_db/blob_db.o
  CC       shared-objects/utilities/blob_db/blob_db_impl.o
  CC       shared-objects/utilities/blob_db/blob_file.o
  CC       shared-objects/utilities/blob_db/blob_log_reader.o
  CC       shared-objects/utilities/blob_db/blob_log_writer.o
  CC       shared-objects/utilities/blob_db/blob_log_format.o
  CC       shared-objects/utilities/blob_db/ttl_extractor.o
  CC       shared-objects/utilities/cassandra/cassandra_compaction_filter.o
  CC       shared-objects/utilities/cassandra/format.o
  CC       shared-objects/utilities/cassandra/merge_operator.o
  CC       shared-objects/utilities/checkpoint/checkpoint_impl.o
  CC       shared-objects/utilities/compaction_filters/remove_emptyvalue_compactionfilter.o
  CC       shared-objects/utilities/convenience/info_log_finder.o
  CC       shared-objects/utilities/date_tiered/date_tiered_db_impl.o
  CC       shared-objects/utilities/debug.o
  CC       shared-objects/utilities/document/document_db.o
  CC       shared-objects/utilities/document/json_document.o
  CC       shared-objects/utilities/document/json_document_builder.o
  CC       shared-objects/utilities/env_mirror.o
  CC       shared-objects/utilities/env_timed.o
  CC       shared-objects/utilities/geodb/geodb_impl.o
  CC       shared-objects/utilities/leveldb_options/leveldb_options.o
  CC       shared-objects/utilities/lua/rocks_lua_compaction_filter.o
  CC       shared-objects/utilities/memory/memory_util.o
  CC       shared-objects/utilities/merge_operators/max.o
  CC       shared-objects/utilities/merge_operators/put.o
  CC       shared-objects/utilities/merge_operators/string_append/stringappend.o
  CC       shared-objects/utilities/merge_operators/string_append/stringappend2.o
  CC       shared-objects/utilities/merge_operators/uint64add.o
  CC       shared-objects/utilities/option_change_migration/option_change_migration.o
  CC       shared-objects/utilities/options/options_util.o
  CC       shared-objects/utilities/persistent_cache/block_cache_tier.o
  CC       shared-objects/utilities/persistent_cache/block_cache_tier_file.o
  CC       shared-objects/utilities/persistent_cache/block_cache_tier_metadata.o
  CC       shared-objects/utilities/persistent_cache/persistent_cache_tier.o
  CC       shared-objects/utilities/persistent_cache/volatile_tier_impl.o
  CC       shared-objects/utilities/redis/redis_lists.o
  CC       shared-objects/utilities/simulator_cache/sim_cache.o
  CC       shared-objects/utilities/spatialdb/spatial_db.o
  CC       shared-objects/utilities/table_properties_collectors/compact_on_deletion_collector.o
  CC       shared-objects/utilities/transactions/optimistic_transaction_db_impl.o
  CC       shared-objects/utilities/transactions/optimistic_transaction.o
  CC       shared-objects/utilities/transactions/pessimistic_transaction.o
  CC       shared-objects/utilities/transactions/pessimistic_transaction_db.o
  CC       shared-objects/utilities/transactions/snapshot_checker.o
  CC       shared-objects/utilities/transactions/transaction_base.o
  CC       shared-objects/utilities/transactions/transaction_db_mutex_impl.o
  CC       shared-objects/utilities/transactions/transaction_lock_mgr.o
  CC       shared-objects/utilities/transactions/transaction_util.o
  CC       shared-objects/utilities/transactions/write_prepared_txn.o
  CC       shared-objects/utilities/ttl/db_ttl_impl.o
  CC       shared-objects/utilities/write_batch_with_index/write_batch_with_index.o
  CC       shared-objects/utilities/write_batch_with_index/write_batch_with_index_internal.o
  CC       shared-objects/tools/ldb_cmd.o
  CC       shared-objects/tools/ldb_tool.o
  CC       shared-objects/tools/sst_dump_tool.o
  CC       shared-objects/utilities/blob_db/blob_dump_tool.o
g++ -Wl,--no-as-needed -shared -Wl,-soname -Wl,librocksdb.so.5.8  -fno-rtti -g -W -Wextra -Wall -Wsign-compare -Wshadow -Wno-unused-parameter -Werror -I. -I./include -std=c++11  -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_FALLOCATE_PRESENT -DZLIB -DROCKSDB_MALLOC_USABLE_SIZE -DROCKSDB_PTHREAD_ADAPTIVE_MUTEX -DROCKSDB_BACKTRACE -DROCKSDB_RANGESYNC_PRESENT -DROCKSDB_SCHED_GETCPU_PRESENT -march=native  -DHAVE_SSE42 -DROCKSDB_SUPPORT_THREAD_LOCAL  -isystem ./third-party/gtest-1.7.0/fused-src -O2 -fno-omit-frame-pointer -momit-leaf-frame-pointer -DNDEBUG -Woverloaded-virtual -Wnon-virtual-dtor -Wno-missing-field-initializers -fPIC shared-objects/cache/clock_cache.o shared-objects/cache/lru_cache.o shared-objects/cache/sharded_cache.o shared-objects/db/builder.o shared-objects/db/c.o shared-objects/db/column_family.o shared-objects/db/compacted_db_impl.o shared-objects/db/compaction.o shared-objects/db/compaction_iterator.o shared-objects/db/compaction_job.o shared-objects/db/compaction_picker.o shared-objects/db/compaction_picker_universal.o shared-objects/db/convenience.o shared-objects/db/db_filesnapshot.o shared-objects/db/db_impl.o shared-objects/db/db_impl_write.o shared-objects/db/db_impl_compaction_flush.o shared-objects/db/db_impl_files.o shared-objects/db/db_impl_open.o shared-objects/db/db_impl_debug.o shared-objects/db/db_impl_experimental.o shared-objects/db/db_impl_readonly.o shared-objects/db/db_info_dumper.o shared-objects/db/db_iter.o shared-objects/db/dbformat.o shared-objects/db/event_helpers.o shared-objects/db/experimental.o shared-objects/db/external_sst_file_ingestion_job.o shared-objects/db/file_indexer.o shared-objects/db/flush_job.o shared-objects/db/flush_scheduler.o shared-objects/db/forward_iterator.o shared-objects/db/internal_stats.o shared-objects/db/log_reader.o shared-objects/db/log_writer.o shared-objects/db/malloc_stats.o shared-objects/db/managed_iterator.o shared-objects/db/memtable.o shared-objects/db/memtable_list.o shared-objects/db/merge_helper.o shared-objects/db/merge_operator.o shared-objects/db/range_del_aggregator.o shared-objects/db/repair.o shared-objects/db/snapshot_impl.o shared-objects/db/table_cache.o shared-objects/db/table_properties_collector.o shared-objects/db/transaction_log_impl.o shared-objects/db/version_builder.o shared-objects/db/version_edit.o shared-objects/db/version_set.o shared-objects/db/wal_manager.o shared-objects/db/write_batch.o shared-objects/db/write_batch_base.o shared-objects/db/write_controller.o shared-objects/db/write_thread.o shared-objects/env/env.o shared-objects/env/env_chroot.o shared-objects/env/env_encryption.o shared-objects/env/env_hdfs.o shared-objects/env/env_posix.o shared-objects/env/io_posix.o shared-objects/env/mock_env.o shared-objects/memtable/alloc_tracker.o shared-objects/memtable/hash_cuckoo_rep.o shared-objects/memtable/hash_linklist_rep.o shared-objects/memtable/hash_skiplist_rep.o shared-objects/memtable/skiplistrep.o shared-objects/memtable/vectorrep.o shared-objects/memtable/write_buffer_manager.o shared-objects/monitoring/histogram.o shared-objects/monitoring/histogram_windowing.o shared-objects/monitoring/instrumented_mutex.o shared-objects/monitoring/iostats_context.o shared-objects/monitoring/perf_context.o shared-objects/monitoring/perf_level.o shared-objects/monitoring/statistics.o shared-objects/monitoring/thread_status_impl.o shared-objects/monitoring/thread_status_updater.o shared-objects/monitoring/thread_status_updater_debug.o shared-objects/monitoring/thread_status_util.o shared-objects/monitoring/thread_status_util_debug.o shared-objects/options/cf_options.o shared-objects/options/db_options.o shared-objects/options/options.o shared-objects/options/options_helper.o shared-objects/options/options_parser.o shared-objects/options/options_sanity_check.o shared-objects/port/port_posix.o shared-objects/port/stack_trace.o shared-objects/table/adaptive_table_factory.o shared-objects/table/block.o shared-objects/table/block_based_filter_block.o shared-objects/table/block_based_table_builder.o shared-objects/table/block_based_table_factory.o shared-objects/table/block_based_table_reader.o shared-objects/table/block_builder.o shared-objects/table/block_prefix_index.o shared-objects/table/bloom_block.o shared-objects/table/cuckoo_table_builder.o shared-objects/table/cuckoo_table_factory.o shared-objects/table/cuckoo_table_reader.o shared-objects/table/flush_block_policy.o shared-objects/table/format.o shared-objects/table/full_filter_block.o shared-objects/table/get_context.o shared-objects/table/index_builder.o shared-objects/table/iterator.o shared-objects/table/merging_iterator.o shared-objects/table/meta_blocks.o shared-objects/table/partitioned_filter_block.o shared-objects/table/persistent_cache_helper.o shared-objects/table/plain_table_builder.o shared-objects/table/plain_table_factory.o shared-objects/table/plain_table_index.o shared-objects/table/plain_table_key_coding.o shared-objects/table/plain_table_reader.o shared-objects/table/sst_file_writer.o shared-objects/table/table_properties.o shared-objects/table/two_level_iterator.o shared-objects/tools/dump/db_dump_tool.o shared-objects/util/arena.o shared-objects/util/auto_roll_logger.o shared-objects/util/bloom.o shared-objects/util/build_version.o shared-objects/util/coding.o shared-objects/util/compaction_job_stats_impl.o shared-objects/util/comparator.o shared-objects/util/concurrent_arena.o shared-objects/util/crc32c.o shared-objects/util/delete_scheduler.o shared-objects/util/dynamic_bloom.o shared-objects/util/event_logger.o shared-objects/util/file_reader_writer.o shared-objects/util/file_util.o shared-objects/util/filename.o shared-objects/util/filter_policy.o shared-objects/util/hash.o shared-objects/util/log_buffer.o shared-objects/util/murmurhash.o shared-objects/util/random.o shared-objects/util/rate_limiter.o shared-objects/util/slice.o shared-objects/util/sst_file_manager_impl.o shared-objects/util/status.o shared-objects/util/status_message.o shared-objects/util/string_util.o shared-objects/util/sync_point.o shared-objects/util/thread_local.o shared-objects/util/threadpool_imp.o shared-objects/util/transaction_test_util.o shared-objects/util/xxhash.o shared-objects/utilities/backupable/backupable_db.o shared-objects/utilities/blob_db/blob_db.o shared-objects/utilities/blob_db/blob_db_impl.o shared-objects/utilities/blob_db/blob_file.o shared-objects/utilities/blob_db/blob_log_reader.o shared-objects/utilities/blob_db/blob_log_writer.o shared-objects/utilities/blob_db/blob_log_format.o shared-objects/utilities/blob_db/ttl_extractor.o shared-objects/utilities/cassandra/cassandra_compaction_filter.o shared-objects/utilities/cassandra/format.o shared-objects/utilities/cassandra/merge_operator.o shared-objects/utilities/checkpoint/checkpoint_impl.o shared-objects/utilities/compaction_filters/remove_emptyvalue_compactionfilter.o shared-objects/utilities/convenience/info_log_finder.o shared-objects/utilities/date_tiered/date_tiered_db_impl.o shared-objects/utilities/debug.o shared-objects/utilities/document/document_db.o shared-objects/utilities/document/json_document.o shared-objects/utilities/document/json_document_builder.o shared-objects/utilities/env_mirror.o shared-objects/utilities/env_timed.o shared-objects/utilities/geodb/geodb_impl.o shared-objects/utilities/leveldb_options/leveldb_options.o shared-objects/utilities/lua/rocks_lua_compaction_filter.o shared-objects/utilities/memory/memory_util.o shared-objects/utilities/merge_operators/max.o shared-objects/utilities/merge_operators/put.o shared-objects/utilities/merge_operators/string_append/stringappend.o shared-objects/utilities/merge_operators/string_append/stringappend2.o shared-objects/utilities/merge_operators/uint64add.o shared-objects/utilities/option_change_migration/option_change_migration.o shared-objects/utilities/options/options_util.o shared-objects/utilities/persistent_cache/block_cache_tier.o shared-objects/utilities/persistent_cache/block_cache_tier_file.o shared-objects/utilities/persistent_cache/block_cache_tier_metadata.o shared-objects/utilities/persistent_cache/persistent_cache_tier.o shared-objects/utilities/persistent_cache/volatile_tier_impl.o shared-objects/utilities/redis/redis_lists.o shared-objects/utilities/simulator_cache/sim_cache.o shared-objects/utilities/spatialdb/spatial_db.o shared-objects/utilities/table_properties_collectors/compact_on_deletion_collector.o shared-objects/utilities/transactions/optimistic_transaction_db_impl.o shared-objects/utilities/transactions/optimistic_transaction.o shared-objects/utilities/transactions/pessimistic_transaction.o shared-objects/utilities/transactions/pessimistic_transaction_db.o shared-objects/utilities/transactions/snapshot_checker.o shared-objects/utilities/transactions/transaction_base.o shared-objects/utilities/transactions/transaction_db_mutex_impl.o shared-objects/utilities/transactions/transaction_lock_mgr.o shared-objects/utilities/transactions/transaction_util.o shared-objects/utilities/transactions/write_prepared_txn.o shared-objects/utilities/ttl/db_ttl_impl.o shared-objects/utilities/write_batch_with_index/write_batch_with_index.o shared-objects/utilities/write_batch_with_index/write_batch_with_index_internal.o shared-objects/tools/ldb_cmd.o shared-objects/tools/ldb_tool.o shared-objects/tools/sst_dump_tool.o shared-objects/utilities/blob_db/blob_dump_tool.o  -lpthread -lrt -lz -o librocksdb.so.5.8.0
ln -fs librocksdb.so.5.8.0 librocksdb.so
ln -fs librocksdb.so.5.8.0 librocksdb.so.5
ln -fs librocksdb.so.5.8.0 librocksdb.so.5.8
make: 警告：检测到时钟错误。您的创建可能是不完整的。
```

Install
```
[vagrant@localhost rocksdb]$ sudo make install shared_lib
Makefile:127: Warning: Compiling in debug mode. Don't use the resulting binary in production
make: Warning: File 'make_config.mk' has modification time 674 s in the future
  GEN      util/build_version.cc
install -d /usr/local/lib
for header_dir in `find "include/rocksdb" -type d`; do \
	install -d /usr/local/$header_dir; \
done
for header in `find "include/rocksdb" -type f -name *.h`; do \
	install -C -m 644 $header /usr/local/$header; \
done
install -C -m 755 librocksdb_debug.a /usr/local/lib
[ -e librocksdb_debug.so.5.8.0 ] && make install-shared || :
g++ -Wl,--no-as-needed -shared -Wl,-soname -Wl,librocksdb_debug.so.5.8  -DROCKSDB_USE_RTTI -g -W -Wextra -Wall -Wsign-compare -Wshadow -Wno-unused-parameter -Werror -I. -I./include -std=c++11  -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_FALLOCATE_PRESENT -DZLIB -DROCKSDB_MALLOC_USABLE_SIZE -DROCKSDB_PTHREAD_ADAPTIVE_MUTEX -DROCKSDB_BACKTRACE -DROCKSDB_RANGESYNC_PRESENT -DROCKSDB_SCHED_GETCPU_PRESENT -march=native  -DHAVE_SSE42 -DROCKSDB_SUPPORT_THREAD_LOCAL  -isystem ./third-party/gtest-1.7.0/fused-src -O2 -fno-omit-frame-pointer -momit-leaf-frame-pointer -Woverloaded-virtual -Wnon-virtual-dtor -Wno-missing-field-initializers -fPIC shared-objects/cache/clock_cache.o shared-objects/cache/lru_cache.o shared-objects/cache/sharded_cache.o shared-objects/db/builder.o shared-objects/db/c.o shared-objects/db/column_family.o shared-objects/db/compacted_db_impl.o shared-objects/db/compaction.o shared-objects/db/compaction_iterator.o shared-objects/db/compaction_job.o shared-objects/db/compaction_picker.o shared-objects/db/compaction_picker_universal.o shared-objects/db/convenience.o shared-objects/db/db_filesnapshot.o shared-objects/db/db_impl.o shared-objects/db/db_impl_write.o shared-objects/db/db_impl_compaction_flush.o shared-objects/db/db_impl_files.o shared-objects/db/db_impl_open.o shared-objects/db/db_impl_debug.o shared-objects/db/db_impl_experimental.o shared-objects/db/db_impl_readonly.o shared-objects/db/db_info_dumper.o shared-objects/db/db_iter.o shared-objects/db/dbformat.o shared-objects/db/event_helpers.o shared-objects/db/experimental.o shared-objects/db/external_sst_file_ingestion_job.o shared-objects/db/file_indexer.o shared-objects/db/flush_job.o shared-objects/db/flush_scheduler.o shared-objects/db/forward_iterator.o shared-objects/db/internal_stats.o shared-objects/db/log_reader.o shared-objects/db/log_writer.o shared-objects/db/malloc_stats.o shared-objects/db/managed_iterator.o shared-objects/db/memtable.o shared-objects/db/memtable_list.o shared-objects/db/merge_helper.o shared-objects/db/merge_operator.o shared-objects/db/range_del_aggregator.o shared-objects/db/repair.o shared-objects/db/snapshot_impl.o shared-objects/db/table_cache.o shared-objects/db/table_properties_collector.o shared-objects/db/transaction_log_impl.o shared-objects/db/version_builder.o shared-objects/db/version_edit.o shared-objects/db/version_set.o shared-objects/db/wal_manager.o shared-objects/db/write_batch.o shared-objects/db/write_batch_base.o shared-objects/db/write_controller.o shared-objects/db/write_thread.o shared-objects/env/env.o shared-objects/env/env_chroot.o shared-objects/env/env_encryption.o shared-objects/env/env_hdfs.o shared-objects/env/env_posix.o shared-objects/env/io_posix.o shared-objects/env/mock_env.o shared-objects/memtable/alloc_tracker.o shared-objects/memtable/hash_cuckoo_rep.o shared-objects/memtable/hash_linklist_rep.o shared-objects/memtable/hash_skiplist_rep.o shared-objects/memtable/skiplistrep.o shared-objects/memtable/vectorrep.o shared-objects/memtable/write_buffer_manager.o shared-objects/monitoring/histogram.o shared-objects/monitoring/histogram_windowing.o shared-objects/monitoring/instrumented_mutex.o shared-objects/monitoring/iostats_context.o shared-objects/monitoring/perf_context.o shared-objects/monitoring/perf_level.o shared-objects/monitoring/statistics.o shared-objects/monitoring/thread_status_impl.o shared-objects/monitoring/thread_status_updater.o shared-objects/monitoring/thread_status_updater_debug.o shared-objects/monitoring/thread_status_util.o shared-objects/monitoring/thread_status_util_debug.o shared-objects/options/cf_options.o shared-objects/options/db_options.o shared-objects/options/options.o shared-objects/options/options_helper.o shared-objects/options/options_parser.o shared-objects/options/options_sanity_check.o shared-objects/port/port_posix.o shared-objects/port/stack_trace.o shared-objects/table/adaptive_table_factory.o shared-objects/table/block.o shared-objects/table/block_based_filter_block.o shared-objects/table/block_based_table_builder.o shared-objects/table/block_based_table_factory.o shared-objects/table/block_based_table_reader.o shared-objects/table/block_builder.o shared-objects/table/block_prefix_index.o shared-objects/table/bloom_block.o shared-objects/table/cuckoo_table_builder.o shared-objects/table/cuckoo_table_factory.o shared-objects/table/cuckoo_table_reader.o shared-objects/table/flush_block_policy.o shared-objects/table/format.o shared-objects/table/full_filter_block.o shared-objects/table/get_context.o shared-objects/table/index_builder.o shared-objects/table/iterator.o shared-objects/table/merging_iterator.o shared-objects/table/meta_blocks.o shared-objects/table/partitioned_filter_block.o shared-objects/table/persistent_cache_helper.o shared-objects/table/plain_table_builder.o shared-objects/table/plain_table_factory.o shared-objects/table/plain_table_index.o shared-objects/table/plain_table_key_coding.o shared-objects/table/plain_table_reader.o shared-objects/table/sst_file_writer.o shared-objects/table/table_properties.o shared-objects/table/two_level_iterator.o shared-objects/tools/dump/db_dump_tool.o shared-objects/util/arena.o shared-objects/util/auto_roll_logger.o shared-objects/util/bloom.o shared-objects/util/build_version.o shared-objects/util/coding.o shared-objects/util/compaction_job_stats_impl.o shared-objects/util/comparator.o shared-objects/util/concurrent_arena.o shared-objects/util/crc32c.o shared-objects/util/delete_scheduler.o shared-objects/util/dynamic_bloom.o shared-objects/util/event_logger.o shared-objects/util/file_reader_writer.o shared-objects/util/file_util.o shared-objects/util/filename.o shared-objects/util/filter_policy.o shared-objects/util/hash.o shared-objects/util/log_buffer.o shared-objects/util/murmurhash.o shared-objects/util/random.o shared-objects/util/rate_limiter.o shared-objects/util/slice.o shared-objects/util/sst_file_manager_impl.o shared-objects/util/status.o shared-objects/util/status_message.o shared-objects/util/string_util.o shared-objects/util/sync_point.o shared-objects/util/thread_local.o shared-objects/util/threadpool_imp.o shared-objects/util/transaction_test_util.o shared-objects/util/xxhash.o shared-objects/utilities/backupable/backupable_db.o shared-objects/utilities/blob_db/blob_db.o shared-objects/utilities/blob_db/blob_db_impl.o shared-objects/utilities/blob_db/blob_file.o shared-objects/utilities/blob_db/blob_log_reader.o shared-objects/utilities/blob_db/blob_log_writer.o shared-objects/utilities/blob_db/blob_log_format.o shared-objects/utilities/blob_db/ttl_extractor.o shared-objects/utilities/cassandra/cassandra_compaction_filter.o shared-objects/utilities/cassandra/format.o shared-objects/utilities/cassandra/merge_operator.o shared-objects/utilities/checkpoint/checkpoint_impl.o shared-objects/utilities/compaction_filters/remove_emptyvalue_compactionfilter.o shared-objects/utilities/convenience/info_log_finder.o shared-objects/utilities/date_tiered/date_tiered_db_impl.o shared-objects/utilities/debug.o shared-objects/utilities/document/document_db.o shared-objects/utilities/document/json_document.o shared-objects/utilities/document/json_document_builder.o shared-objects/utilities/env_mirror.o shared-objects/utilities/env_timed.o shared-objects/utilities/geodb/geodb_impl.o shared-objects/utilities/leveldb_options/leveldb_options.o shared-objects/utilities/lua/rocks_lua_compaction_filter.o shared-objects/utilities/memory/memory_util.o shared-objects/utilities/merge_operators/max.o shared-objects/utilities/merge_operators/put.o shared-objects/utilities/merge_operators/string_append/stringappend.o shared-objects/utilities/merge_operators/string_append/stringappend2.o shared-objects/utilities/merge_operators/uint64add.o shared-objects/utilities/option_change_migration/option_change_migration.o shared-objects/utilities/options/options_util.o shared-objects/utilities/persistent_cache/block_cache_tier.o shared-objects/utilities/persistent_cache/block_cache_tier_file.o shared-objects/utilities/persistent_cache/block_cache_tier_metadata.o shared-objects/utilities/persistent_cache/persistent_cache_tier.o shared-objects/utilities/persistent_cache/volatile_tier_impl.o shared-objects/utilities/redis/redis_lists.o shared-objects/utilities/simulator_cache/sim_cache.o shared-objects/utilities/spatialdb/spatial_db.o shared-objects/utilities/table_properties_collectors/compact_on_deletion_collector.o shared-objects/utilities/transactions/optimistic_transaction_db_impl.o shared-objects/utilities/transactions/optimistic_transaction.o shared-objects/utilities/transactions/pessimistic_transaction.o shared-objects/utilities/transactions/pessimistic_transaction_db.o shared-objects/utilities/transactions/snapshot_checker.o shared-objects/utilities/transactions/transaction_base.o shared-objects/utilities/transactions/transaction_db_mutex_impl.o shared-objects/utilities/transactions/transaction_lock_mgr.o shared-objects/utilities/transactions/transaction_util.o shared-objects/utilities/transactions/write_prepared_txn.o shared-objects/utilities/ttl/db_ttl_impl.o shared-objects/utilities/write_batch_with_index/write_batch_with_index.o shared-objects/utilities/write_batch_with_index/write_batch_with_index_internal.o shared-objects/tools/ldb_cmd.o shared-objects/tools/ldb_tool.o shared-objects/tools/sst_dump_tool.o shared-objects/utilities/blob_db/blob_dump_tool.o  -lpthread -lrt -lz -o librocksdb_debug.so.5.8.0
ln -fs librocksdb_debug.so.5.8.0 librocksdb_debug.so
ln -fs librocksdb_debug.so.5.8.0 librocksdb_debug.so.5
ln -fs librocksdb_debug.so.5.8.0 librocksdb_debug.so.5.8
make: 警告：检测到时钟错误。您的创建可能是不完整的。
```

## First Build
Clone
```
[vagrant@localhost github.com]$ git clone --depth=1 https://github.com/facebook/rocksdb facebook/rocksdb
正克隆到 'facebook/rocksdb'...
remote: Counting objects: 935, done.
remote: Compressing objects: 100% (896/896), done.
接收对象中: 100% (935/935), 2.38 MiB | 5.00 KiB/s, 完成.
remote: Total 935 (delta 87), reused 186 (delta 14), pack-reused 0
处理 delta 中: 100% (87/87), 完成.
检查连接... 完成。
```

Revision
```
[vagrant@localhost github.com]$ cd facebook/rocksdb/ && git log
commit 88a2776db52015c7c32767bdd38b6d072fd1dbd6
Author: Islam AbdelRahman <tec@fb.com>
Date:   Mon Jun 20 11:26:25 2016 -0700

    Update SstFileWriter to use bottommost_compression if avaliable
    
    Summary: SstFileWriter ignore Options::bottommost_compression, update it to use bottommost_compression if available
    
    Test Plan:
    make check -j64
    verified used compression using ./sst_dump
    
    Reviewers: sdong
    
    Reviewed By: sdong
    
    Subscribers: andrewkr, dhruba, yoshinorim
    
    Differential Revision: https://reviews.facebook.net/D59841
```

Directory
```
[vagrant@localhost rocksdb]$ ls
appveyor.yml   CMakeLists.txt   DEFAULT_OPTIONS_HISTORY.md  hdfs        java                  memtable   ROCKSDB_LITE.md  thirdparty.inc  utilities
arcanist_util  CONTRIBUTING.md  doc                         HISTORY.md  LANGUAGE-BINDINGS.md  PATENTS    src.mk           tools           Vagrantfile
AUTHORS        coverage         DUMP_FORMAT.md              include     LICENSE               port       table            USERS.md        WINDOWS_PORT.md
build_tools    db               examples                    INSTALL.md  Makefile              README.md  third-party      util
```

Make
```
[vagrant@localhost rocksdb]$ make static_lib
  GEN      util/build_version.cc
  GEN      util/build_version.cc
  CC       db/auto_roll_logger.o
  CC       db/builder.o
  CC       db/c.o
### snippets ###
  CC       tools/ldb_cmd.o
  CC       tools/ldb_tool.o
  CC       tools/sst_dump_tool.o
  AR       librocksdb.a
ar: 正在创建 librocksdb.a
```

Install
```
[vagrant@localhost rocksdb]$ sudo make install
  GEN      util/build_version.cc
install -d /usr/local/lib
for header_dir in `find "include/rocksdb" -type d`; do \
	install -d /usr/local/$header_dir; \
done
for header in `find "include/rocksdb" -type f -name *.h`; do \
	install -C -m 644 $header /usr/local/$header; \
done
install -C -m 755 librocksdb.a /usr/local/lib
[ -e librocksdb.so.4.9.0 ] && make install-shared || :
```

Make Lib
```
[vagrant@localhost rocksdb]$ make shared_lib
  GEN      util/build_version.cc
g++ -Wl,--no-as-needed -shared -Wl,-soname -Wl,librocksdb.so.4.9  -g -W -Wextra -Wall -Wsign-compare -Wshadow -Wno-unused-parameter -I. -I./include -std=c++11  -DROCKSDB_PLATFORM_POSIX -DROCKSDB_LIB_IO_POSIX  -DOS_LINUX -fno-builtin-memcmp -DROCKSDB_FALLOCATE_PRESENT -DZLIB -DROCKSDB_MALLOC_USABLE_SIZE -DROCKSDB_PTHREAD_ADAPTIVE_MUTEX -DROCKSDB_BACKTRACE -march=native   -isystem ./third-party/gtest-1.7.0/fused-src -O2 -fno-omit-frame-pointer -momit-leaf-frame-pointer -DNDEBUG -Woverloaded-virtual -Wnon-virtual-dtor -Wno-missing-field-initializers -fPIC db/auto_roll_logger.cc db/builder.cc db/c.cc db/column_family.cc db/compacted_db_impl.cc db/compaction.cc db/compaction_iterator.cc db/compaction_job.cc db/compaction_picker.cc db/convenience.cc db/db_filesnapshot.cc db/dbformat.cc db/db_impl.cc db/db_impl_debug.cc db/db_impl_readonly.cc db/db_impl_experimental.cc db/db_info_dumper.cc db/db_iter.cc db/experimental.cc db/event_helpers.cc db/file_indexer.cc db/filename.cc db/flush_job.cc db/flush_scheduler.cc db/forward_iterator.cc db/internal_stats.cc db/log_reader.cc db/log_writer.cc db/managed_iterator.cc db/memtable_allocator.cc db/memtable.cc db/memtable_list.cc db/merge_helper.cc db/merge_operator.cc db/repair.cc db/snapshot_impl.cc db/table_cache.cc db/table_properties_collector.cc db/transaction_log_impl.cc db/version_builder.cc db/version_edit.cc db/version_set.cc db/wal_manager.cc db/write_batch.cc db/write_batch_base.cc db/write_controller.cc db/write_thread.cc db/xfunc_test_points.cc memtable/hash_cuckoo_rep.cc memtable/hash_linklist_rep.cc memtable/hash_skiplist_rep.cc memtable/skiplistrep.cc memtable/vectorrep.cc port/stack_trace.cc port/port_posix.cc table/adaptive_table_factory.cc table/block_based_filter_block.cc table/block_based_table_builder.cc table/block_based_table_factory.cc table/block_based_table_reader.cc table/block_builder.cc table/block.cc table/block_prefix_index.cc table/bloom_block.cc table/cuckoo_table_builder.cc table/cuckoo_table_factory.cc table/cuckoo_table_reader.cc table/flush_block_policy.cc table/format.cc table/full_filter_block.cc table/get_context.cc table/iterator.cc table/merger.cc table/meta_blocks.cc table/sst_file_writer.cc table/plain_table_builder.cc table/plain_table_factory.cc table/plain_table_index.cc table/plain_table_key_coding.cc table/plain_table_reader.cc table/persistent_cache_helper.cc table/table_properties.cc table/two_level_iterator.cc tools/dump/db_dump_tool.cc util/arena.cc util/bloom.cc util/build_version.cc util/cache.cc util/coding.cc util/comparator.cc util/compaction_job_stats_impl.cc util/concurrent_arena.cc util/crc32c.cc util/delete_scheduler.cc util/dynamic_bloom.cc util/env.cc util/env_chroot.cc util/env_hdfs.cc util/env_posix.cc util/io_posix.cc util/threadpool.cc util/transaction_test_util.cc util/sst_file_manager_impl.cc util/file_util.cc util/file_reader_writer.cc util/filter_policy.cc util/hash.cc util/histogram.cc util/histogram_windowing.cc util/instrumented_mutex.cc util/iostats_context.cc utilities/backupable/backupable_db.cc utilities/convenience/info_log_finder.cc utilities/checkpoint/checkpoint.cc utilities/compaction_filters/remove_emptyvalue_compactionfilter.cc utilities/document/document_db.cc utilities/document/json_document_builder.cc utilities/document/json_document.cc utilities/env_mirror.cc utilities/env_registry.cc utilities/flashcache/flashcache.cc utilities/geodb/geodb_impl.cc utilities/leveldb_options/leveldb_options.cc utilities/memory/memory_util.cc utilities/merge_operators/put.cc utilities/merge_operators/max.cc utilities/merge_operators/string_append/stringappend2.cc utilities/merge_operators/string_append/stringappend.cc utilities/merge_operators/uint64add.cc utilities/options/options_util.cc utilities/persistent_cache/persistent_cache_tier.cc utilities/persistent_cache/volatile_tier_impl.cc utilities/redis/redis_lists.cc utilities/simulator_cache/sim_cache.cc utilities/spatialdb/spatial_db.cc utilities/table_properties_collectors/compact_on_deletion_collector.cc utilities/transactions/optimistic_transaction_impl.cc utilities/transactions/optimistic_transaction_db_impl.cc utilities/transactions/transaction_base.cc utilities/transactions/transaction_db_impl.cc utilities/transactions/transaction_db_mutex_impl.cc utilities/transactions/transaction_lock_mgr.cc utilities/transactions/transaction_impl.cc utilities/transactions/transaction_util.cc utilities/ttl/db_ttl_impl.cc utilities/write_batch_with_index/write_batch_with_index.cc utilities/write_batch_with_index/write_batch_with_index_internal.cc util/event_logger.cc util/log_buffer.cc util/logging.cc util/memenv.cc util/murmurhash.cc util/mutable_cf_options.cc util/options.cc util/options_helper.cc util/options_parser.cc util/options_sanity_check.cc util/perf_context.cc util/perf_level.cc util/random.cc util/rate_limiter.cc util/slice.cc util/statistics.cc util/status.cc util/status_message.cc util/string_util.cc util/sync_point.cc util/thread_local.cc util/thread_status_impl.cc util/thread_status_updater.cc util/thread_status_updater_debug.cc util/thread_status_util.cc util/thread_status_util_debug.cc util/xfunc.cc util/xxhash.cc  tools/ldb_cmd.cc tools/ldb_tool.cc tools/sst_dump_tool.cc  \
	 -lpthread -lrt -lz -o librocksdb.so.4.9.0
```

Install Lib
```
[vagrant@localhost rocksdb]$ sudo make install_shared
Makefile:101: Warning: Compiling in debug mode. Don't use the resulting binary in production
  GEN      util/build_version.cc
make: *** No rule to make target 'install_shared'。 停止。
[vagrant@localhost rocksdb]$ sudo make install
  GEN      util/build_version.cc
install -d /usr/local/lib
for header_dir in `find "include/rocksdb" -type d`; do \
	install -d /usr/local/$header_dir; \
done
for header in `find "include/rocksdb" -type f -name *.h`; do \
	install -C -m 644 $header /usr/local/$header; \
done
install -C -m 755 librocksdb.a /usr/local/lib
[ -e librocksdb.so.4.9.0 ] && make install-shared || :
make[1]: Entering directory '/data/src/github.com/facebook/rocksdb'
  GEN      util/build_version.cc
install -d /usr/local/lib
for header_dir in `find "include/rocksdb" -type d`; do \
	install -d /usr/local/$header_dir; \
done
for header in `find "include/rocksdb" -type f -name *.h`; do \
	install -C -m 644 $header /usr/local/$header; \
done
install -C -m 755 librocksdb.so.4.9.0 /usr/local/lib && \
	ln -fs librocksdb.so.4.9.0 /usr/local/lib/librocksdb.so.4.9 && \
	ln -fs librocksdb.so.4.9.0 /usr/local/lib/librocksdb.so.4 && \
	ln -fs librocksdb.so.4.9.0 /usr/local/lib/librocksdb.so
make[1]: Leaving directory '/data/src/github.com/facebook/rocksdb'
```
