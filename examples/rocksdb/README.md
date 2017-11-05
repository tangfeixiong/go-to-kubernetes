Test
```
[vagrant@localhost ~]$ LD_LIBRARY_PATH=/Users/fanhongling/Downloads/workspace/src/github.com/facebook/rocksdb ~/go/bin/rocksdb-example
/tmp/rocksdb-data
bar
```

Database
```
[vagrant@localhost ~]$ ls -l /tmp/rocksdb-data/
总用量 152
-rw-r--r--. 1 vagrant vagrant    52 10月 26 18:28 000003.log
-rw-r--r--. 1 vagrant vagrant    16 10月 26 18:28 CURRENT
-rw-r--r--. 1 vagrant vagrant    37 10月 26 18:28 IDENTITY
-rw-r--r--. 1 vagrant vagrant     0 10月 26 18:28 LOCK
-rw-rw-r--. 1 vagrant vagrant 14562 10月 26 18:28 LOG
-rw-r--r--. 1 vagrant vagrant    13 10月 26 18:28 MANIFEST-000001
-rw-r--r--. 1 vagrant vagrant  4235 10月 26 18:28 OPTIONS-000005
```

LOG
```
[vagrant@localhost rocksdb-data]$ cat LOG 
2017/10/26-18:28:17.750019 7f016a1f5a00 RocksDB version: 5.8.0
2017/10/26-18:28:17.750452 7f016a1f5a00 Git sha rocksdb_build_git_sha:731895214bb729f4b748c8051d2d4c81927ff47f
2017/10/26-18:28:17.750457 7f016a1f5a00 Compile date Oct 17 2017
2017/10/26-18:28:17.750637 7f016a1f5a00 DB SUMMARY
2017/10/26-18:28:17.751086 7f016a1f5a00 SST files in /tmp/rocksdb-data dir, Total Num: 0, files: 
2017/10/26-18:28:17.751091 7f016a1f5a00 Write Ahead Log file in /tmp/rocksdb-data: 
2017/10/26-18:28:17.751095 7f016a1f5a00                         Options.error_if_exists: 0
2017/10/26-18:28:17.751271 7f016a1f5a00                       Options.create_if_missing: 1
2017/10/26-18:28:17.751275 7f016a1f5a00                         Options.paranoid_checks: 1
2017/10/26-18:28:17.751277 7f016a1f5a00                                     Options.env: 0x7f0169fdeca0
2017/10/26-18:28:17.751279 7f016a1f5a00                                Options.info_log: 0x2565980
2017/10/26-18:28:17.751280 7f016a1f5a00                Options.max_file_opening_threads: 16
2017/10/26-18:28:17.751282 7f016a1f5a00                              Options.statistics: (nil)
2017/10/26-18:28:17.751284 7f016a1f5a00                               Options.use_fsync: 0
2017/10/26-18:28:17.751286 7f016a1f5a00                       Options.max_log_file_size: 0
2017/10/26-18:28:17.751288 7f016a1f5a00                  Options.max_manifest_file_size: 18446744073709551615
2017/10/26-18:28:17.751290 7f016a1f5a00                   Options.log_file_time_to_roll: 0
2017/10/26-18:28:17.751292 7f016a1f5a00                       Options.keep_log_file_num: 1000
2017/10/26-18:28:17.751294 7f016a1f5a00                    Options.recycle_log_file_num: 0
2017/10/26-18:28:17.751296 7f016a1f5a00                         Options.allow_fallocate: 1
2017/10/26-18:28:17.751298 7f016a1f5a00                        Options.allow_mmap_reads: 0
2017/10/26-18:28:17.751299 7f016a1f5a00                       Options.allow_mmap_writes: 0
2017/10/26-18:28:17.751301 7f016a1f5a00                        Options.use_direct_reads: 0
2017/10/26-18:28:17.751303 7f016a1f5a00                        Options.use_direct_io_for_flush_and_compaction: 0
2017/10/26-18:28:17.751305 7f016a1f5a00          Options.create_missing_column_families: 0
2017/10/26-18:28:17.751307 7f016a1f5a00                              Options.db_log_dir: 
2017/10/26-18:28:17.751309 7f016a1f5a00                                 Options.wal_dir: /tmp/rocksdb-data
2017/10/26-18:28:17.751311 7f016a1f5a00                Options.table_cache_numshardbits: 6
2017/10/26-18:28:17.751313 7f016a1f5a00                      Options.max_subcompactions: 1
2017/10/26-18:28:17.751314 7f016a1f5a00                  Options.max_background_flushes: -1
2017/10/26-18:28:17.751316 7f016a1f5a00                         Options.WAL_ttl_seconds: 0
2017/10/26-18:28:17.751318 7f016a1f5a00                       Options.WAL_size_limit_MB: 0
2017/10/26-18:28:17.751320 7f016a1f5a00             Options.manifest_preallocation_size: 4194304
2017/10/26-18:28:17.751322 7f016a1f5a00                     Options.is_fd_close_on_exec: 1
2017/10/26-18:28:17.751324 7f016a1f5a00                   Options.advise_random_on_open: 1
2017/10/26-18:28:17.751326 7f016a1f5a00                    Options.db_write_buffer_size: 0
2017/10/26-18:28:17.751328 7f016a1f5a00                    Options.write_buffer_manager: 0x25654a0
2017/10/26-18:28:17.751329 7f016a1f5a00         Options.access_hint_on_compaction_start: 1
2017/10/26-18:28:17.751331 7f016a1f5a00  Options.new_table_reader_for_compaction_inputs: 0
2017/10/26-18:28:17.751507 7f016a1f5a00               Options.compaction_readahead_size: 0
2017/10/26-18:28:17.751510 7f016a1f5a00           Options.random_access_max_buffer_size: 1048576
2017/10/26-18:28:17.751512 7f016a1f5a00           Options.writable_file_max_buffer_size: 1048576
2017/10/26-18:28:17.751514 7f016a1f5a00                      Options.use_adaptive_mutex: 0
2017/10/26-18:28:17.751516 7f016a1f5a00                            Options.rate_limiter: (nil)
2017/10/26-18:28:17.751519 7f016a1f5a00     Options.sst_file_manager.rate_bytes_per_sec: 0
2017/10/26-18:28:17.751521 7f016a1f5a00                       Options.wal_recovery_mode: 2
2017/10/26-18:28:17.751528 7f016a1f5a00                  Options.enable_thread_tracking: 0
2017/10/26-18:28:17.751530 7f016a1f5a00                  Options.enable_pipelined_write: 0
2017/10/26-18:28:17.751532 7f016a1f5a00         Options.allow_concurrent_memtable_write: 1
2017/10/26-18:28:17.751534 7f016a1f5a00      Options.enable_write_thread_adaptive_yield: 1
2017/10/26-18:28:17.751536 7f016a1f5a00             Options.write_thread_max_yield_usec: 100
2017/10/26-18:28:17.751538 7f016a1f5a00            Options.write_thread_slow_yield_usec: 3
2017/10/26-18:28:17.751539 7f016a1f5a00                               Options.row_cache: None
2017/10/26-18:28:17.751541 7f016a1f5a00                              Options.wal_filter: None
2017/10/26-18:28:17.751543 7f016a1f5a00             Options.avoid_flush_during_recovery: 0
2017/10/26-18:28:17.751545 7f016a1f5a00             Options.allow_ingest_behind: 0
2017/10/26-18:28:17.751547 7f016a1f5a00             Options.concurrent_prepare: 0
2017/10/26-18:28:17.751549 7f016a1f5a00             Options.manual_wal_flush: 0
2017/10/26-18:28:17.751551 7f016a1f5a00             Options.seq_per_batch: 0
2017/10/26-18:28:17.751554 7f016a1f5a00             Options.max_background_jobs: 2
2017/10/26-18:28:17.751556 7f016a1f5a00             Options.max_background_compactions: -1
2017/10/26-18:28:17.751558 7f016a1f5a00             Options.avoid_flush_during_shutdown: 0
2017/10/26-18:28:17.751559 7f016a1f5a00             Options.delayed_write_rate : 16777216
2017/10/26-18:28:17.751561 7f016a1f5a00             Options.max_total_wal_size: 0
2017/10/26-18:28:17.751563 7f016a1f5a00             Options.delete_obsolete_files_period_micros: 21600000000
2017/10/26-18:28:17.751565 7f016a1f5a00                   Options.stats_dump_period_sec: 600
2017/10/26-18:28:17.751567 7f016a1f5a00                          Options.max_open_files: -1
2017/10/26-18:28:17.751569 7f016a1f5a00                          Options.bytes_per_sync: 0
2017/10/26-18:28:17.751571 7f016a1f5a00                      Options.wal_bytes_per_sync: 0
2017/10/26-18:28:17.751573 7f016a1f5a00 Compression algorithms supported:
2017/10/26-18:28:17.751575 7f016a1f5a00 	Snappy supported: 0
2017/10/26-18:28:17.751577 7f016a1f5a00 	Zlib supported: 1
2017/10/26-18:28:17.751579 7f016a1f5a00 	Bzip supported: 0
2017/10/26-18:28:17.751581 7f016a1f5a00 	LZ4 supported: 0
2017/10/26-18:28:17.751583 7f016a1f5a00 	ZSTDNotFinal supported: 0
2017/10/26-18:28:17.751585 7f016a1f5a00 	ZSTD supported: 0
2017/10/26-18:28:17.751590 7f016a1f5a00 Fast CRC32 supported: Supported on x86Not supported on x86
2017/10/26-18:28:17.752626 7f016a1f5a00 [db/db_impl_open.cc:216] Creating manifest 1 
2017/10/26-18:28:17.759142 7f016a1f5a00 [db/version_set.cc:2740] Recovering from manifest file: MANIFEST-000001
2017/10/26-18:28:17.760191 7f016a1f5a00 [db/column_family.cc:413] --------------- Options for column family [default]:
2017/10/26-18:28:17.760375 7f016a1f5a00               Options.comparator: leveldb.BytewiseComparator
2017/10/26-18:28:17.760380 7f016a1f5a00           Options.merge_operator: None
2017/10/26-18:28:17.760382 7f016a1f5a00        Options.compaction_filter: None
2017/10/26-18:28:17.760384 7f016a1f5a00        Options.compaction_filter_factory: None
2017/10/26-18:28:17.760386 7f016a1f5a00         Options.memtable_factory: SkipListFactory
2017/10/26-18:28:17.760388 7f016a1f5a00            Options.table_factory: BlockBasedTable
2017/10/26-18:28:17.760422 7f016a1f5a00            table_factory options:   flush_block_policy_factory: FlushBlockBySizePolicyFactory (0x255aad0)
  cache_index_and_filter_blocks: 0
  cache_index_and_filter_blocks_with_high_priority: 0
  pin_l0_filter_and_index_blocks_in_cache: 0
  index_type: 0
  hash_index_allow_collision: 1
  checksum: 1
  no_block_cache: 0
  block_cache: 0x255d760
  block_cache_name: LRUCache
  block_cache_options:
    capacity : 3221225472
    num_shard_bits : 6
    strict_capacity_limit : 0
    high_pri_pool_ratio: 0.000
  block_cache_compressed: (nil)
  persistent_cache: (nil)
  block_size: 4096
  block_size_deviation: 10
  block_restart_interval: 16
  index_block_restart_interval: 1
  metadata_block_size: 4096
  partition_filters: 0
  use_delta_encoding: 1
  filter_policy: nullptr
  whole_key_filtering: 1
  verify_compression: 0
  read_amp_bytes_per_bit: 0
  format_version: 2
2017/10/26-18:28:17.760432 7f016a1f5a00        Options.write_buffer_size: 67108864
2017/10/26-18:28:17.760434 7f016a1f5a00  Options.max_write_buffer_number: 2
2017/10/26-18:28:17.760438 7f016a1f5a00          Options.compression: NoCompression
2017/10/26-18:28:17.760440 7f016a1f5a00                  Options.bottommost_compression: Disabled
2017/10/26-18:28:17.760442 7f016a1f5a00       Options.prefix_extractor: nullptr
2017/10/26-18:28:17.760444 7f016a1f5a00   Options.memtable_insert_with_hint_prefix_extractor: nullptr
2017/10/26-18:28:17.760446 7f016a1f5a00             Options.num_levels: 7
2017/10/26-18:28:17.760448 7f016a1f5a00        Options.min_write_buffer_number_to_merge: 1
2017/10/26-18:28:17.760449 7f016a1f5a00     Options.max_write_buffer_number_to_maintain: 0
2017/10/26-18:28:17.760451 7f016a1f5a00            Options.compression_opts.window_bits: -14
2017/10/26-18:28:17.760453 7f016a1f5a00                  Options.compression_opts.level: -1
2017/10/26-18:28:17.760455 7f016a1f5a00               Options.compression_opts.strategy: 0
2017/10/26-18:28:17.760457 7f016a1f5a00         Options.compression_opts.max_dict_bytes: 0
2017/10/26-18:28:17.760459 7f016a1f5a00      Options.level0_file_num_compaction_trigger: 4
2017/10/26-18:28:17.760461 7f016a1f5a00          Options.level0_slowdown_writes_trigger: 20
2017/10/26-18:28:17.760463 7f016a1f5a00              Options.level0_stop_writes_trigger: 36
2017/10/26-18:28:17.760465 7f016a1f5a00                   Options.target_file_size_base: 67108864
2017/10/26-18:28:17.760466 7f016a1f5a00             Options.target_file_size_multiplier: 1
2017/10/26-18:28:17.760468 7f016a1f5a00                Options.max_bytes_for_level_base: 268435456
2017/10/26-18:28:17.760470 7f016a1f5a00 Options.level_compaction_dynamic_level_bytes: 0
2017/10/26-18:28:17.760472 7f016a1f5a00          Options.max_bytes_for_level_multiplier: 10.000000
2017/10/26-18:28:17.760476 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[0]: 1
2017/10/26-18:28:17.760478 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[1]: 1
2017/10/26-18:28:17.760480 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[2]: 1
2017/10/26-18:28:17.760482 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[3]: 1
2017/10/26-18:28:17.760483 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[4]: 1
2017/10/26-18:28:17.760485 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[5]: 1
2017/10/26-18:28:17.760487 7f016a1f5a00 Options.max_bytes_for_level_multiplier_addtl[6]: 1
2017/10/26-18:28:17.760489 7f016a1f5a00       Options.max_sequential_skip_in_iterations: 8
2017/10/26-18:28:17.760491 7f016a1f5a00                    Options.max_compaction_bytes: 1677721600
2017/10/26-18:28:17.760493 7f016a1f5a00                        Options.arena_block_size: 8388608
2017/10/26-18:28:17.760495 7f016a1f5a00   Options.soft_pending_compaction_bytes_limit: 68719476736
2017/10/26-18:28:17.760497 7f016a1f5a00   Options.hard_pending_compaction_bytes_limit: 274877906944
2017/10/26-18:28:17.760499 7f016a1f5a00       Options.rate_limit_delay_max_milliseconds: 100
2017/10/26-18:28:17.760501 7f016a1f5a00                Options.disable_auto_compactions: 0
2017/10/26-18:28:17.760504 7f016a1f5a00                        Options.compaction_style: kCompactionStyleLevel
2017/10/26-18:28:17.760506 7f016a1f5a00                          Options.compaction_pri: kByCompensatedSize
2017/10/26-18:28:17.760508 7f016a1f5a00 Options.compaction_options_universal.size_ratio: 1
2017/10/26-18:28:17.760510 7f016a1f5a00 Options.compaction_options_universal.min_merge_width: 2
2017/10/26-18:28:17.760512 7f016a1f5a00 Options.compaction_options_universal.max_merge_width: 4294967295
2017/10/26-18:28:17.760513 7f016a1f5a00 Options.compaction_options_universal.max_size_amplification_percent: 200
2017/10/26-18:28:17.760518 7f016a1f5a00 Options.compaction_options_universal.compression_size_percent: -1
2017/10/26-18:28:17.760520 7f016a1f5a00 Options.compaction_options_universal.stop_style: kCompactionStopStyleTotalSize
2017/10/26-18:28:17.760522 7f016a1f5a00 Options.compaction_options_fifo.max_table_files_size: 1073741824
2017/10/26-18:28:17.760524 7f016a1f5a00 Options.compaction_options_fifo.allow_compaction: 0
2017/10/26-18:28:17.760526 7f016a1f5a00 Options.compaction_options_fifo.ttl: 0
2017/10/26-18:28:17.760528 7f016a1f5a00                   Options.table_properties_collectors: 
2017/10/26-18:28:17.760530 7f016a1f5a00                   Options.inplace_update_support: 0
2017/10/26-18:28:17.760532 7f016a1f5a00                 Options.inplace_update_num_locks: 10000
2017/10/26-18:28:17.760534 7f016a1f5a00               Options.memtable_prefix_bloom_size_ratio: 0.000000
2017/10/26-18:28:17.760536 7f016a1f5a00   Options.memtable_huge_page_size: 0
2017/10/26-18:28:17.760538 7f016a1f5a00                           Options.bloom_locality: 0
2017/10/26-18:28:17.760540 7f016a1f5a00                    Options.max_successive_merges: 0
2017/10/26-18:28:17.760542 7f016a1f5a00                Options.optimize_filters_for_hits: 0
2017/10/26-18:28:17.760544 7f016a1f5a00                Options.paranoid_file_checks: 0
2017/10/26-18:28:17.760546 7f016a1f5a00                Options.force_consistency_checks: 0
2017/10/26-18:28:17.760548 7f016a1f5a00                Options.report_bg_io_stats: 0
2017/10/26-18:28:17.767676 7f016a1f5a00 [db/version_set.cc:3004] Recovered from manifest file:/tmp/rocksdb-data/MANIFEST-000001 succeeded,manifest_file_number is 1, next_file_number is 3, last_sequence is 0, log_number is 0,prev_log_number is 0,max_column_family is 0
2017/10/26-18:28:17.767688 7f016a1f5a00 [db/version_set.cc:3012] Column family [default] (ID 0), log number is 0
2017/10/26-18:28:17.775850 7f016a1f5a00 [db/db_impl_open.cc:1135] DB pointer 0x2564830
2017/10/26-18:28:17.790614 7f016a1f5a00 [db/db_impl.cc:227] Shutdown: canceling all background work
2017/10/26-18:28:17.796806 7f016a1f5a00 [db/db_impl.cc:357] Shutdown complete
```

OPTIONS
```
[vagrant@localhost rocksdb-data]$ head OPTIONS-000005 
# This is a RocksDB option file.
#
# For detailed file format spec, please refer to the example file
# in examples/rocksdb_option_file_example.ini
#

[Version]
  rocksdb_version=5.8.0
  options_file_version=1.1

[vagrant@localhost rocksdb-data]$ cat OPTIONS-000005 
# This is a RocksDB option file.
#
# For detailed file format spec, please refer to the example file
# in examples/rocksdb_option_file_example.ini
#

[Version]
  rocksdb_version=5.8.0
  options_file_version=1.1

[DBOptions]
  seq_per_batch=false
  concurrent_prepare=false
  avoid_flush_during_shutdown=false
  info_log_level=INFO_LEVEL
  access_hint_on_compaction_start=NORMAL
  write_thread_max_yield_usec=100
  enable_write_thread_adaptive_yield=true
  write_thread_slow_yield_usec=3
  fail_if_options_file_error=false
  wal_recovery_mode=kPointInTimeRecovery
  max_manifest_file_size=18446744073709551615
  delete_obsolete_files_period_micros=21600000000
  WAL_ttl_seconds=0
  WAL_size_limit_MB=0
  max_subcompactions=1
  wal_dir=/tmp/rocksdb-data
  dump_malloc_stats=false
  db_log_dir=
  recycle_log_file_num=0
  keep_log_file_num=1000
  enable_pipelined_write=false
  delayed_write_rate=16777216
  db_write_buffer_size=0
  table_cache_numshardbits=6
  avoid_flush_during_recovery=false
  max_open_files=-1
  max_file_opening_threads=16
  max_background_flushes=-1
  log_file_time_to_roll=0
  base_background_compactions=-1
  max_background_compactions=-1
  use_fsync=false
  allow_concurrent_memtable_write=true
  writable_file_max_buffer_size=1048576
  random_access_max_buffer_size=1048576
  new_table_reader_for_compaction_inputs=false
  max_background_jobs=2
  skip_log_error_on_recovery=false
  paranoid_checks=true
  max_total_wal_size=0
  is_fd_close_on_exec=true
  allow_ingest_behind=false
  error_if_exists=false
  wal_bytes_per_sync=0
  stats_dump_period_sec=600
  create_missing_column_families=false
  manual_wal_flush=false
  create_if_missing=true
  allow_2pc=false
  skip_stats_update_on_db_open=false
  use_direct_io_for_flush_and_compaction=false
  bytes_per_sync=0
  max_log_file_size=0
  manifest_preallocation_size=4194304
  use_direct_reads=false
  allow_mmap_writes=false
  allow_fallocate=true
  compaction_readahead_size=0
  allow_mmap_reads=false
  use_adaptive_mutex=false
  enable_thread_tracking=false
  advise_random_on_open=true
  

[CFOptions "default"]
  compaction_pri=kByCompensatedSize
  compaction_filter_factory=nullptr
  memtable_factory=SkipListFactory
  bottommost_compression=kDisableCompressionOption
  compression=kNoCompression
  max_sequential_skip_in_iterations=8
  max_bytes_for_level_multiplier_additional=1:1:1:1:1:1:1
  compression_per_level=
  max_bytes_for_level_multiplier=10.000000
  max_bytes_for_level_base=268435456
  table_factory=BlockBasedTable
  max_successive_merges=0
  arena_block_size=8388608
  merge_operator=nullptr
  target_file_size_multiplier=1
  num_levels=7
  min_write_buffer_number_to_merge=1
  prefix_extractor=nullptr
  bloom_locality=0
  max_write_buffer_number=2
  level0_stop_writes_trigger=36
  level0_slowdown_writes_trigger=20
  level0_file_num_compaction_trigger=4
  write_buffer_size=67108864
  memtable_huge_page_size=0
  max_compaction_bytes=1677721600
  hard_pending_compaction_bytes_limit=274877906944
  target_file_size_base=67108864
  soft_pending_compaction_bytes_limit=68719476736
  comparator=leveldb.BytewiseComparator
  memtable_insert_with_hint_prefix_extractor=nullptr
  force_consistency_checks=false
  max_write_buffer_number_to_maintain=0
  paranoid_file_checks=false
  optimize_filters_for_hits=false
  level_compaction_dynamic_level_bytes=false
  inplace_update_support=false
  compaction_style=kCompactionStyleLevel
  compaction_filter=nullptr
  disable_auto_compactions=false
  inplace_update_num_locks=10000
  memtable_prefix_bloom_size_ratio=0.000000
  report_bg_io_stats=false
  
[TableOptions/BlockBasedTable "default"]
  read_amp_bytes_per_bit=8589934592
  format_version=2
  whole_key_filtering=true
  filter_policy=nullptr
  verify_compression=false
  block_size_deviation=10
  block_size=4096
  partition_filters=false
  checksum=kCRC32c
  hash_index_allow_collision=true
  index_block_restart_interval=1
  block_restart_interval=16
  no_block_cache=false
  pin_l0_filter_and_index_blocks_in_cache=false
  cache_index_and_filter_blocks_with_high_priority=false
  metadata_block_size=4096
  cache_index_and_filter_blocks=false
  index_type=kBinarySearch
  flush_block_policy_factory=FlushBlockBySizePolicyFactory
```