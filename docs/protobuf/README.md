Build Protobuf development environment
======================================

Prerequistes
------------

Ubuntu 14.04 VBox

Get repository
--------------

With `git`
```
vagrant@vagrant-ubuntu-trusty-64:~$ git clone https://github.com/google/protobuf /work/src/github.com/google/protobuf

vagrant@vagrant-ubuntu-trusty-64:~$ cd /work/src/github.com/google/protobuf/
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ cat src/README.md 
Protocol Buffers - Google's data interchange format
===================================================

[![Build Status](https://travis-ci.org/google/protobuf.svg?branch=master)](https://travis-ci.org/google/protobuf) [![Build status](https://ci.appveyor.com/api/projects/status/73ctee6ua4w2ruin?svg=true)](https://ci.appveyor.com/project/protobuf/protobuf)

Copyright 2008 Google Inc.

https://developers.google.com/protocol-buffers/

C++ Installation - Unix
-----------------------

To build protobuf from source, the following tools are needed:

  * autoconf
  * automake
  * libtool
  * curl (used to download gmock)
  * make
  * g++
  * unzip

On Ubuntu, you can install them with:

  $ sudo apt-get install autoconf automake libtool curl make g++ unzip

On other platforms, please use the corresponding package managing tool to
install them before proceeding.

If you get the source from github, you need to generate the configure script
first:

    $ ./autogen.sh

This will download gmock source (which is used for C++ Protocol Buffer
unit-tests) to the current directory and run automake, autoconf, etc.
to generate the configure script and various template makefiles.

You can skip this step if you are using a release package (which already
contains gmock and the configure script).

To build and install the C++ Protocol Buffer runtime and the Protocol
Buffer compiler (protoc) execute the following:

    $ ./configure
    $ make
    $ make check
    $ sudo make install
    $ sudo ldconfig # refresh shared library cache.

If "make check" fails, you can still install, but it is likely that
some features of this library will not work correctly on your system.
Proceed at your own risk.

For advanced usage information on configure and make, please refer to the
autoconf documentation:

  http://www.gnu.org/software/autoconf/manual/autoconf.html#Running-configure-Scripts


Usage
-----

The complete documentation for Protocol Buffers is available via the
web at:

    https://developers.google.com/protocol-buffers/
```

Install _Ubuntu_ build tools
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf/src$ sudo apt-get install autoconf automake libtool curl make g++ unzip
Reading package lists... Done
Building dependency tree       
Reading state information... Done
g++ is already the newest version.
g++ set to manually installed.
make is already the newest version.
make set to manually installed.
curl is already the newest version.
unzip is already the newest version.
unzip set to manually installed.
The following extra packages will be installed:
  autotools-dev libltdl-dev m4
Suggested packages:
  autoconf2.13 autoconf-archive gnu-standards autoconf-doc gettext libtool-doc
  automaken gfortran fortran95-compiler gcj-jdk
The following NEW packages will be installed:
  autoconf automake autotools-dev libltdl-dev libtool m4
0 upgraded, 6 newly installed, 0 to remove and 25 not upgraded.
Need to get 1,417 kB of archives.
After this operation, 6,395 kB of additional disk space will be used.
Do you want to continue? [Y/n] y

```

Generate 'makefile'
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ ./autogen.sh 
+ autoreconf -f -i -Wall,no-obsolete
libtoolize: putting auxiliary files in AC_CONFIG_AUX_DIR, `build-aux'.
libtoolize: copying file `build-aux/ltmain.sh'
libtoolize: putting macros in AC_CONFIG_MACRO_DIR, `m4'.
libtoolize: copying file `m4/libtool.m4'
libtoolize: copying file `m4/ltoptions.m4'
libtoolize: copying file `m4/ltsugar.m4'
libtoolize: copying file `m4/ltversion.m4'
libtoolize: copying file `m4/lt~obsolete.m4'
libtoolize: putting auxiliary files in AC_CONFIG_AUX_DIR, `build-aux'.
libtoolize: copying file `build-aux/ltmain.sh'
libtoolize: Consider adding `AC_CONFIG_MACRO_DIR([m4])' to configure.ac and
libtoolize: rerunning libtoolize, to keep the correct libtool macros in-tree.
libtoolize: Consider adding `-I m4' to ACLOCAL_AMFLAGS in Makefile.am.
libtoolize: putting auxiliary files in `.'.
libtoolize: copying file `./ltmain.sh'
libtoolize: putting macros in AC_CONFIG_MACRO_DIR, `m4'.
libtoolize: copying file `m4/libtool.m4'
libtoolize: copying file `m4/ltoptions.m4'
libtoolize: copying file `m4/ltsugar.m4'
libtoolize: copying file `m4/ltversion.m4'
libtoolize: copying file `m4/lt~obsolete.m4'
+ rm -rf autom4te.cache config.h.in~
+ exit 0
```

Optional, clean legacy version
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ make clean

make[3]: Entering directory `/work/src/github.com/google/protobuf/gmock/gtest'
 rm -f samples/sample1_unittest samples/sample10_unittest test/gtest_all_test test/fused_gtest_test
test -z "core" || rm -f core
test -z "lib/libgtest.la lib/libgtest_main.la" || rm -f lib/libgtest.la lib/libgtest_main.la
rm -f lib/so_locations
rm -rf .libs _libs
rm -rf lib/.libs lib/_libs
rm -rf samples/.libs samples/_libs
rm -rf src/.libs src/_libs
rm -rf test/.libs test/_libs
test -z "samples/libsamples.la" || rm -f samples/libsamples.la
rm -f samples/so_locations
rm -f *.o
rm -f fused-src/gtest/*.o
rm -f samples/*.o
rm -f samples/*.lo
rm -f src/*.o
rm -f src/*.lo
rm -f test/*.o
test -z "samples/sample1_unittest.log samples/sample10_unittest.log test/gtest_all_test.log test/fused_gtest_test.log" || rm -f samples/sample1_unittest.log samples/sample10_unittest.log test/gtest_all_test.log test/fused_gtest_test.log
test -z "samples/sample1_unittest.trs samples/sample10_unittest.trs test/gtest_all_test.trs test/fused_gtest_test.trs" || rm -f samples/sample1_unittest.trs samples/sample10_unittest.trs test/gtest_all_test.trs test/fused_gtest_test.trs
test -z "test-suite.log" || rm -f test-suite.log
rm -f *.lo
make[3]: Leaving directory `/work/src/github.com/google/protobuf/gmock/gtest'
make[3]: Entering directory `/work/src/github.com/google/protobuf/gmock'
 rm -f test/gmock-spec-builders_test test/gmock_link_test test/gmock_fused_test
test -z "core" || rm -f core
test -z "lib/libgmock.la lib/libgmock_main.la" || rm -f lib/libgmock.la lib/libgmock_main.la
rm -f lib/so_locations
rm -rf .libs _libs
rm -rf lib/.libs lib/_libs
rm -rf src/.libs src/_libs
rm -rf test/.libs test/_libs
rm -f *.o
rm -f fused-src/*.o
rm -f src/*.o
rm -f src/*.lo
rm -f test/*.o
test -z "test/gmock-spec-builders_test.log test/gmock_link_test.log test/gmock_fused_test.log" || rm -f test/gmock-spec-builders_test.log test/gmock_link_test.log test/gmock_fused_test.log
test -z "test/gmock-spec-builders_test.trs test/gmock_link_test.trs test/gmock_fused_test.trs" || rm -f test/gmock-spec-builders_test.trs test/gmock_link_test.trs test/gmock_fused_test.trs
test -z "test-suite.log" || rm -f test-suite.log
rm -f *.lo
make[3]: Leaving directory `/work/src/github.com/google/protobuf/gmock'
make[2]: Leaving directory `/work/src/github.com/google/protobuf/gmock'
rm -f *.lo
make[1]: Leaving directory `/work/src/github.com/google/protobuf'
Making clean in src
make[1]: Entering directory `/work/src/github.com/google/protobuf/src'
 rm -f protoc
 rm -f protoc protobuf-test protobuf-lazy-descriptor-test protobuf-lite-test test_plugin protobuf-lite-arena-test no-warning-test
test -z "google/protobuf/map_lite_unittest.pb.cc google/protobuf/map_lite_unittest.pb.h google/protobuf/unittest_lite.pb.cc google/protobuf/unittest_lite.pb.h google/protobuf/unittest_no_arena_lite.pb.cc google/protobuf/unittest_no_arena_lite.pb.h google/protobuf/unittest_import_lite.pb.cc google/protobuf/unittest_import_lite.pb.h google/protobuf/unittest_import_public_lite.pb.cc google/protobuf/unittest_import_public_lite.pb.h google/protobuf/any_test.pb.cc google/protobuf/any_test.pb.h google/protobuf/compiler/cpp/cpp_test_bad_identifiers.pb.cc google/protobuf/compiler/cpp/cpp_test_bad_identifiers.pb.h google/protobuf/compiler/cpp/cpp_test_large_enum_value.pb.cc google/protobuf/compiler/cpp/cpp_test_large_enum_value.pb.h google/protobuf/map_proto2_unittest.pb.cc google/protobuf/map_proto2_unittest.pb.h google/protobuf/map_unittest.pb.cc google/protobuf/map_unittest.pb.h google/protobuf/unittest_arena.pb.cc google/protobuf/unittest_arena.pb.h google/protobuf/unittest_custom_options.pb.cc google/protobuf/unittest_custom_options.pb.h google/protobuf/unittest_drop_unknown_fields.pb.cc google/protobuf/unittest_drop_unknown_fields.pb.h google/protobuf/unittest_embed_optimize_for.pb.cc google/protobuf/unittest_embed_optimize_for.pb.h google/protobuf/unittest_empty.pb.cc google/protobuf/unittest_empty.pb.h google/protobuf/unittest_enormous_descriptor.pb.cc google/protobuf/unittest_enormous_descriptor.pb.h google/protobuf/unittest_import.pb.cc google/protobuf/unittest_import.pb.h google/protobuf/unittest_import_public.pb.cc google/protobuf/unittest_import_public.pb.h google/protobuf/unittest_lazy_dependencies.pb.cc google/protobuf/unittest_lazy_dependencies.pb.h google/protobuf/unittest_lazy_dependencies_custom_option.pb.cc google/protobuf/unittest_lazy_dependencies_custom_option.pb.h google/protobuf/unittest_lazy_dependencies_enum.pb.cc google/protobuf/unittest_lazy_dependencies_enum.pb.h google/protobuf/unittest_lite_imports_nonlite.pb.cc google/protobuf/unittest_lite_imports_nonlite.pb.h google/protobuf/unittest_mset.pb.cc google/protobuf/unittest_mset.pb.h google/protobuf/unittest_mset_wire_format.pb.cc google/protobuf/unittest_mset_wire_format.pb.h google/protobuf/unittest_no_arena_import.pb.cc google/protobuf/unittest_no_arena_import.pb.h google/protobuf/unittest_no_arena.pb.cc google/protobuf/unittest_no_arena.pb.h google/protobuf/unittest_no_field_presence.pb.cc google/protobuf/unittest_no_field_presence.pb.h google/protobuf/unittest_no_generic_services.pb.cc google/protobuf/unittest_no_generic_services.pb.h google/protobuf/unittest_optimize_for.pb.cc google/protobuf/unittest_optimize_for.pb.h google/protobuf/unittest.pb.cc google/protobuf/unittest.pb.h google/protobuf/unittest_preserve_unknown_enum2.pb.cc google/protobuf/unittest_preserve_unknown_enum2.pb.h google/protobuf/unittest_preserve_unknown_enum.pb.cc google/protobuf/unittest_preserve_unknown_enum.pb.h google/protobuf/unittest_proto3_arena.pb.cc google/protobuf/unittest_proto3_arena.pb.h google/protobuf/unittest_proto3_arena_lite.pb.cc google/protobuf/unittest_proto3_arena_lite.pb.h google/protobuf/unittest_proto3_lite.pb.cc google/protobuf/unittest_proto3_lite.pb.h google/protobuf/unittest_well_known_types.pb.cc google/protobuf/unittest_well_known_types.pb.h google/protobuf/util/internal/testdata/anys.pb.cc google/protobuf/util/internal/testdata/anys.pb.h google/protobuf/util/internal/testdata/books.pb.cc google/protobuf/util/internal/testdata/books.pb.h google/protobuf/util/internal/testdata/default_value.pb.cc google/protobuf/util/internal/testdata/default_value.pb.h google/protobuf/util/internal/testdata/default_value_test.pb.cc google/protobuf/util/internal/testdata/default_value_test.pb.h google/protobuf/util/internal/testdata/field_mask.pb.cc google/protobuf/util/internal/testdata/field_mask.pb.h google/protobuf/util/internal/testdata/maps.pb.cc google/protobuf/util/internal/testdata/maps.pb.h google/protobuf/util/internal/testdata/oneofs.pb.cc google/protobuf/util/internal/testdata/oneofs.pb.h google/protobuf/util/internal/testdata/proto3.pb.cc google/protobuf/util/internal/testdata/proto3.pb.h google/protobuf/util/internal/testdata/struct.pb.cc google/protobuf/util/internal/testdata/struct.pb.h google/protobuf/util/internal/testdata/timestamp_duration.pb.cc google/protobuf/util/internal/testdata/timestamp_duration.pb.h google/protobuf/util/internal/testdata/wrappers.pb.cc google/protobuf/util/internal/testdata/wrappers.pb.h google/protobuf/util/json_format_proto3.pb.cc google/protobuf/util/json_format_proto3.pb.h google/protobuf/util/message_differencer_unittest.pb.cc google/protobuf/util/message_differencer_unittest.pb.h unittest_proto_middleman testzip.jar testzip.list testzip.proto testzip.zip no_warning_test.cc google/protobuf/compiler/js/well_known_types_embed.cc js_embed" || rm -f google/protobuf/map_lite_unittest.pb.cc google/protobuf/map_lite_unittest.pb.h google/protobuf/unittest_lite.pb.cc google/protobuf/unittest_lite.pb.h google/protobuf/unittest_no_arena_lite.pb.cc google/protobuf/unittest_no_arena_lite.pb.h google/protobuf/unittest_import_lite.pb.cc google/protobuf/unittest_import_lite.pb.h google/protobuf/unittest_import_public_lite.pb.cc google/protobuf/unittest_import_public_lite.pb.h google/protobuf/any_test.pb.cc google/protobuf/any_test.pb.h google/protobuf/compiler/cpp/cpp_test_bad_identifiers.pb.cc google/protobuf/compiler/cpp/cpp_test_bad_identifiers.pb.h google/protobuf/compiler/cpp/cpp_test_large_enum_value.pb.cc google/protobuf/compiler/cpp/cpp_test_large_enum_value.pb.h google/protobuf/map_proto2_unittest.pb.cc google/protobuf/map_proto2_unittest.pb.h google/protobuf/map_unittest.pb.cc google/protobuf/map_unittest.pb.h google/protobuf/unittest_arena.pb.cc google/protobuf/unittest_arena.pb.h google/protobuf/unittest_custom_options.pb.cc google/protobuf/unittest_custom_options.pb.h google/protobuf/unittest_drop_unknown_fields.pb.cc google/protobuf/unittest_drop_unknown_fields.pb.h google/protobuf/unittest_embed_optimize_for.pb.cc google/protobuf/unittest_embed_optimize_for.pb.h google/protobuf/unittest_empty.pb.cc google/protobuf/unittest_empty.pb.h google/protobuf/unittest_enormous_descriptor.pb.cc google/protobuf/unittest_enormous_descriptor.pb.h google/protobuf/unittest_import.pb.cc google/protobuf/unittest_import.pb.h google/protobuf/unittest_import_public.pb.cc google/protobuf/unittest_import_public.pb.h google/protobuf/unittest_lazy_dependencies.pb.cc google/protobuf/unittest_lazy_dependencies.pb.h google/protobuf/unittest_lazy_dependencies_custom_option.pb.cc google/protobuf/unittest_lazy_dependencies_custom_option.pb.h google/protobuf/unittest_lazy_dependencies_enum.pb.cc google/protobuf/unittest_lazy_dependencies_enum.pb.h google/protobuf/unittest_lite_imports_nonlite.pb.cc google/protobuf/unittest_lite_imports_nonlite.pb.h google/protobuf/unittest_mset.pb.cc google/protobuf/unittest_mset.pb.h google/protobuf/unittest_mset_wire_format.pb.cc google/protobuf/unittest_mset_wire_format.pb.h google/protobuf/unittest_no_arena_import.pb.cc google/protobuf/unittest_no_arena_import.pb.h google/protobuf/unittest_no_arena.pb.cc google/protobuf/unittest_no_arena.pb.h google/protobuf/unittest_no_field_presence.pb.cc google/protobuf/unittest_no_field_presence.pb.h google/protobuf/unittest_no_generic_services.pb.cc google/protobuf/unittest_no_generic_services.pb.h google/protobuf/unittest_optimize_for.pb.cc google/protobuf/unittest_optimize_for.pb.h google/protobuf/unittest.pb.cc google/protobuf/unittest.pb.h google/protobuf/unittest_preserve_unknown_enum2.pb.cc google/protobuf/unittest_preserve_unknown_enum2.pb.h google/protobuf/unittest_preserve_unknown_enum.pb.cc google/protobuf/unittest_preserve_unknown_enum.pb.h google/protobuf/unittest_proto3_arena.pb.cc google/protobuf/unittest_proto3_arena.pb.h google/protobuf/unittest_proto3_arena_lite.pb.cc google/protobuf/unittest_proto3_arena_lite.pb.h google/protobuf/unittest_proto3_lite.pb.cc google/protobuf/unittest_proto3_lite.pb.h google/protobuf/unittest_well_known_types.pb.cc google/protobuf/unittest_well_known_types.pb.h google/protobuf/util/internal/testdata/anys.pb.cc google/protobuf/util/internal/testdata/anys.pb.h google/protobuf/util/internal/testdata/books.pb.cc google/protobuf/util/internal/testdata/books.pb.h google/protobuf/util/internal/testdata/default_value.pb.cc google/protobuf/util/internal/testdata/default_value.pb.h google/protobuf/util/internal/testdata/default_value_test.pb.cc google/protobuf/util/internal/testdata/default_value_test.pb.h google/protobuf/util/internal/testdata/field_mask.pb.cc google/protobuf/util/internal/testdata/field_mask.pb.h google/protobuf/util/internal/testdata/maps.pb.cc google/protobuf/util/internal/testdata/maps.pb.h google/protobuf/util/internal/testdata/oneofs.pb.cc google/protobuf/util/internal/testdata/oneofs.pb.h google/protobuf/util/internal/testdata/proto3.pb.cc google/protobuf/util/internal/testdata/proto3.pb.h google/protobuf/util/internal/testdata/struct.pb.cc google/protobuf/util/internal/testdata/struct.pb.h google/protobuf/util/internal/testdata/timestamp_duration.pb.cc google/protobuf/util/internal/testdata/timestamp_duration.pb.h google/protobuf/util/internal/testdata/wrappers.pb.cc google/protobuf/util/internal/testdata/wrappers.pb.h google/protobuf/util/json_format_proto3.pb.cc google/protobuf/util/json_format_proto3.pb.h google/protobuf/util/message_differencer_unittest.pb.cc google/protobuf/util/message_differencer_unittest.pb.h unittest_proto_middleman testzip.jar testzip.list testzip.proto testzip.zip no_warning_test.cc google/protobuf/compiler/js/well_known_types_embed.cc js_embed
test -z "libprotobuf-lite.la libprotobuf.la libprotoc.la" || rm -f libprotobuf-lite.la libprotobuf.la libprotoc.la
rm -f ./so_locations
rm -rf .libs _libs
rm -rf google/protobuf/.libs google/protobuf/_libs
rm -rf google/protobuf/compiler/.libs google/protobuf/compiler/_libs
rm -rf google/protobuf/compiler/cpp/.libs google/protobuf/compiler/cpp/_libs
rm -rf google/protobuf/compiler/csharp/.libs google/protobuf/compiler/csharp/_libs
rm -rf google/protobuf/compiler/java/.libs google/protobuf/compiler/java/_libs
rm -rf google/protobuf/compiler/javanano/.libs google/protobuf/compiler/javanano/_libs
rm -rf google/protobuf/compiler/js/.libs google/protobuf/compiler/js/_libs
rm -rf google/protobuf/compiler/objectivec/.libs google/protobuf/compiler/objectivec/_libs
rm -rf google/protobuf/compiler/php/.libs google/protobuf/compiler/php/_libs
rm -rf google/protobuf/compiler/python/.libs google/protobuf/compiler/python/_libs
rm -rf google/protobuf/compiler/ruby/.libs google/protobuf/compiler/ruby/_libs
rm -rf google/protobuf/io/.libs google/protobuf/io/_libs
rm -rf google/protobuf/stubs/.libs google/protobuf/stubs/_libs
rm -rf google/protobuf/util/.libs google/protobuf/util/_libs
rm -rf google/protobuf/util/internal/.libs google/protobuf/util/internal/_libs
rm -f *.loT
rm -f *.o
rm -f google/protobuf/*.o
rm -f google/protobuf/*.lo
rm -f google/protobuf/compiler/*.o
rm -f google/protobuf/compiler/*.lo
rm -f google/protobuf/compiler/cpp/*.o
rm -f google/protobuf/compiler/cpp/*.lo
rm -f google/protobuf/compiler/csharp/*.o
rm -f google/protobuf/compiler/csharp/*.lo
rm -f google/protobuf/compiler/java/*.o
rm -f google/protobuf/compiler/java/*.lo
rm -f google/protobuf/compiler/javanano/*.o
rm -f google/protobuf/compiler/javanano/*.lo
rm -f google/protobuf/compiler/js/*.o
rm -f google/protobuf/compiler/js/*.lo
rm -f google/protobuf/compiler/objectivec/*.o
rm -f google/protobuf/compiler/objectivec/*.lo
rm -f google/protobuf/compiler/php/*.o
rm -f google/protobuf/compiler/php/*.lo
rm -f google/protobuf/compiler/python/*.o
rm -f google/protobuf/compiler/python/*.lo
rm -f google/protobuf/compiler/ruby/*.o
rm -f google/protobuf/compiler/ruby/*.lo
rm -f google/protobuf/io/*.o
rm -f google/protobuf/io/*.lo
rm -f google/protobuf/stubs/*.o
rm -f google/protobuf/stubs/*.lo
rm -f google/protobuf/testing/*.o
rm -f google/protobuf/util/*.o
rm -f google/protobuf/util/*.lo
rm -f google/protobuf/util/internal/*.o
rm -f google/protobuf/util/internal/*.lo
rm -f google/protobuf/util/internal/testdata/*.o
test -z "protobuf-test.log protobuf-lazy-descriptor-test.log protobuf-lite-test.log google/protobuf/compiler/zip_output_unittest.sh.log protobuf-lite-arena-test.log no-warning-test.log" || rm -f protobuf-test.log protobuf-lazy-descriptor-test.log protobuf-lite-test.log google/protobuf/compiler/zip_output_unittest.sh.log protobuf-lite-arena-test.log no-warning-test.log
test -z "protobuf-test.trs protobuf-lazy-descriptor-test.trs protobuf-lite-test.trs google/protobuf/compiler/zip_output_unittest.sh.trs protobuf-lite-arena-test.trs no-warning-test.trs" || rm -f protobuf-test.trs protobuf-lazy-descriptor-test.trs protobuf-lite-test.trs google/protobuf/compiler/zip_output_unittest.sh.trs protobuf-lite-arena-test.trs no-warning-test.trs
test -z "test-suite.log" || rm -f test-suite.log
rm -f *.lo
make[1]: Leaving directory `/work/src/github.com/google/protobuf/src'
```

Build
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ make
make  all-recursive
make[1]: Entering directory `/work/src/github.com/google/protobuf'
Making all in .
make[2]: Entering directory `/work/src/github.com/google/protobuf'
make[2]: Nothing to be done for `all-am'.
make[2]: Leaving directory `/work/src/github.com/google/protobuf'
Making all in src
make[2]: Entering directory `/work/src/github.com/google/protobuf/src'
depbase=`echo google/protobuf/compiler/main.o | sed 's|[^/]*$|.deps/&|;s|\.o$||'`;\
	g++ -std=c++11 -DHAVE_CONFIG_H -I. -I..    -pthread -DHAVE_PTHREAD=1  -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/compiler/main.o -MD -MP -MF $depbase.Tpo -c -o google/protobuf/compiler/main.o google/protobuf/compiler/main.cc &&\
	mv -f $depbase.Tpo $depbase.Po
depbase=`echo google/protobuf/stubs/atomicops_internals_x86_gcc.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
	/bin/sh ../libtool  --tag=CXX   --mode=compile g++ -std=c++11 -DHAVE_CONFIG_H -I. -I..    -pthread -DHAVE_PTHREAD=1  -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/stubs/atomicops_internals_x86_gcc.lo -MD -MP -MF $depbase.Tpo -c -o google/protobuf/stubs/atomicops_internals_x86_gcc.lo google/protobuf/stubs/atomicops_internals_x86_gcc.cc &&\
	mv -f $depbase.Tpo $depbase.Plo

/bin/sh ../libtool  --tag=CXX   --mode=link g++ -std=c++11 -pthread -DHAVE_PTHREAD=1  -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 13:0:0 -export-dynamic -no-undefined   -o libprotoc.la -rpath /usr/local/lib google/protobuf/compiler/code_generator.lo google/protobuf/compiler/command_line_interface.lo google/protobuf/compiler/plugin.lo google/protobuf/compiler/plugin.pb.lo google/protobuf/compiler/profile.pb.lo google/protobuf/compiler/subprocess.lo google/protobuf/compiler/zip_writer.lo google/protobuf/compiler/cpp/cpp_enum.lo google/protobuf/compiler/cpp/cpp_enum_field.lo google/protobuf/compiler/cpp/cpp_extension.lo google/protobuf/compiler/cpp/cpp_field.lo google/protobuf/compiler/cpp/cpp_file.lo google/protobuf/compiler/cpp/cpp_generator.lo google/protobuf/compiler/cpp/cpp_helpers.lo google/protobuf/compiler/cpp/cpp_map_field.lo google/protobuf/compiler/cpp/cpp_message.lo google/protobuf/compiler/cpp/cpp_message_field.lo google/protobuf/compiler/cpp/cpp_primitive_field.lo google/protobuf/compiler/cpp/cpp_service.lo google/protobuf/compiler/cpp/cpp_string_field.lo google/protobuf/compiler/java/java_context.lo google/protobuf/compiler/java/java_enum.lo google/protobuf/compiler/java/java_enum_lite.lo google/protobuf/compiler/java/java_enum_field.lo google/protobuf/compiler/java/java_enum_field_lite.lo google/protobuf/compiler/java/java_extension.lo google/protobuf/compiler/java/java_extension_lite.lo google/protobuf/compiler/java/java_field.lo google/protobuf/compiler/java/java_file.lo google/protobuf/compiler/java/java_generator.lo google/protobuf/compiler/java/java_generator_factory.lo google/protobuf/compiler/java/java_helpers.lo google/protobuf/compiler/java/java_lazy_message_field.lo google/protobuf/compiler/java/java_lazy_message_field_lite.lo google/protobuf/compiler/java/java_map_field.lo google/protobuf/compiler/java/java_map_field_lite.lo google/protobuf/compiler/java/java_message.lo google/protobuf/compiler/java/java_message_lite.lo google/protobuf/compiler/java/java_message_builder.lo google/protobuf/compiler/java/java_message_builder_lite.lo google/protobuf/compiler/java/java_message_field.lo google/protobuf/compiler/java/java_message_field_lite.lo google/protobuf/compiler/java/java_name_resolver.lo google/protobuf/compiler/java/java_primitive_field.lo google/protobuf/compiler/java/java_primitive_field_lite.lo google/protobuf/compiler/java/java_shared_code_generator.lo google/protobuf/compiler/java/java_service.lo google/protobuf/compiler/java/java_string_field.lo google/protobuf/compiler/java/java_string_field_lite.lo google/protobuf/compiler/java/java_doc_comment.lo google/protobuf/compiler/js/js_generator.lo google/protobuf/compiler/js/well_known_types_embed.lo google/protobuf/compiler/javanano/javanano_enum.lo google/protobuf/compiler/javanano/javanano_enum_field.lo google/protobuf/compiler/javanano/javanano_extension.lo google/protobuf/compiler/javanano/javanano_field.lo google/protobuf/compiler/javanano/javanano_file.lo google/protobuf/compiler/javanano/javanano_generator.lo google/protobuf/compiler/javanano/javanano_helpers.lo google/protobuf/compiler/javanano/javanano_map_field.lo google/protobuf/compiler/javanano/javanano_message.lo google/protobuf/compiler/javanano/javanano_message_field.lo google/protobuf/compiler/javanano/javanano_primitive_field.lo google/protobuf/compiler/objectivec/objectivec_enum.lo google/protobuf/compiler/objectivec/objectivec_enum_field.lo google/protobuf/compiler/objectivec/objectivec_extension.lo google/protobuf/compiler/objectivec/objectivec_field.lo google/protobuf/compiler/objectivec/objectivec_file.lo google/protobuf/compiler/objectivec/objectivec_generator.lo google/protobuf/compiler/objectivec/objectivec_helpers.lo google/protobuf/compiler/objectivec/objectivec_map_field.lo google/protobuf/compiler/objectivec/objectivec_message.lo google/protobuf/compiler/objectivec/objectivec_message_field.lo google/protobuf/compiler/objectivec/objectivec_oneof.lo google/protobuf/compiler/objectivec/objectivec_primitive_field.lo google/protobuf/compiler/php/php_generator.lo google/protobuf/compiler/python/python_generator.lo google/protobuf/compiler/ruby/ruby_generator.lo google/protobuf/compiler/csharp/csharp_doc_comment.lo google/protobuf/compiler/csharp/csharp_enum.lo google/protobuf/compiler/csharp/csharp_enum_field.lo google/protobuf/compiler/csharp/csharp_field_base.lo google/protobuf/compiler/csharp/csharp_generator.lo google/protobuf/compiler/csharp/csharp_helpers.lo google/protobuf/compiler/csharp/csharp_map_field.lo google/protobuf/compiler/csharp/csharp_message.lo google/protobuf/compiler/csharp/csharp_message_field.lo google/protobuf/compiler/csharp/csharp_primitive_field.lo google/protobuf/compiler/csharp/csharp_reflection_class.lo google/protobuf/compiler/csharp/csharp_repeated_enum_field.lo google/protobuf/compiler/csharp/csharp_repeated_message_field.lo google/protobuf/compiler/csharp/csharp_repeated_primitive_field.lo google/protobuf/compiler/csharp/csharp_source_generator_base.lo google/protobuf/compiler/csharp/csharp_wrapper_field.lo -lpthread libprotobuf.la 
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crti.o /usr/lib/gcc/x86_64-linux-gnu/4.8/crtbeginS.o  google/protobuf/compiler/.libs/code_generator.o google/protobuf/compiler/.libs/command_line_interface.o google/protobuf/compiler/.libs/plugin.o google/protobuf/compiler/.libs/plugin.pb.o google/protobuf/compiler/.libs/profile.pb.o google/protobuf/compiler/.libs/subprocess.o google/protobuf/compiler/.libs/zip_writer.o google/protobuf/compiler/cpp/.libs/cpp_enum.o google/protobuf/compiler/cpp/.libs/cpp_enum_field.o google/protobuf/compiler/cpp/.libs/cpp_extension.o google/protobuf/compiler/cpp/.libs/cpp_field.o google/protobuf/compiler/cpp/.libs/cpp_file.o google/protobuf/compiler/cpp/.libs/cpp_generator.o google/protobuf/compiler/cpp/.libs/cpp_helpers.o google/protobuf/compiler/cpp/.libs/cpp_map_field.o google/protobuf/compiler/cpp/.libs/cpp_message.o google/protobuf/compiler/cpp/.libs/cpp_message_field.o google/protobuf/compiler/cpp/.libs/cpp_primitive_field.o google/protobuf/compiler/cpp/.libs/cpp_service.o google/protobuf/compiler/cpp/.libs/cpp_string_field.o google/protobuf/compiler/java/.libs/java_context.o google/protobuf/compiler/java/.libs/java_enum.o google/protobuf/compiler/java/.libs/java_enum_lite.o google/protobuf/compiler/java/.libs/java_enum_field.o google/protobuf/compiler/java/.libs/java_enum_field_lite.o google/protobuf/compiler/java/.libs/java_extension.o google/protobuf/compiler/java/.libs/java_extension_lite.o google/protobuf/compiler/java/.libs/java_field.o google/protobuf/compiler/java/.libs/java_file.o google/protobuf/compiler/java/.libs/java_generator.o google/protobuf/compiler/java/.libs/java_generator_factory.o google/protobuf/compiler/java/.libs/java_helpers.o google/protobuf/compiler/java/.libs/java_lazy_message_field.o google/protobuf/compiler/java/.libs/java_lazy_message_field_lite.o google/protobuf/compiler/java/.libs/java_map_field.o google/protobuf/compiler/java/.libs/java_map_field_lite.o google/protobuf/compiler/java/.libs/java_message.o google/protobuf/compiler/java/.libs/java_message_lite.o google/protobuf/compiler/java/.libs/java_message_builder.o google/protobuf/compiler/java/.libs/java_message_builder_lite.o google/protobuf/compiler/java/.libs/java_message_field.o google/protobuf/compiler/java/.libs/java_message_field_lite.o google/protobuf/compiler/java/.libs/java_name_resolver.o google/protobuf/compiler/java/.libs/java_primitive_field.o google/protobuf/compiler/java/.libs/java_primitive_field_lite.o google/protobuf/compiler/java/.libs/java_shared_code_generator.o google/protobuf/compiler/java/.libs/java_service.o google/protobuf/compiler/java/.libs/java_string_field.o google/protobuf/compiler/java/.libs/java_string_field_lite.o google/protobuf/compiler/java/.libs/java_doc_comment.o google/protobuf/compiler/js/.libs/js_generator.o google/protobuf/compiler/js/.libs/well_known_types_embed.o google/protobuf/compiler/javanano/.libs/javanano_enum.o google/protobuf/compiler/javanano/.libs/javanano_enum_field.o google/protobuf/compiler/javanano/.libs/javanano_extension.o google/protobuf/compiler/javanano/.libs/javanano_field.o google/protobuf/compiler/javanano/.libs/javanano_file.o google/protobuf/compiler/javanano/.libs/javanano_generator.o google/protobuf/compiler/javanano/.libs/javanano_helpers.o google/protobuf/compiler/javanano/.libs/javanano_map_field.o google/protobuf/compiler/javanano/.libs/javanano_message.o google/protobuf/compiler/javanano/.libs/javanano_message_field.o google/protobuf/compiler/javanano/.libs/javanano_primitive_field.o google/protobuf/compiler/objectivec/.libs/objectivec_enum.o google/protobuf/compiler/objectivec/.libs/objectivec_enum_field.o google/protobuf/compiler/objectivec/.libs/objectivec_extension.o google/protobuf/compiler/objectivec/.libs/objectivec_field.o google/protobuf/compiler/objectivec/.libs/objectivec_file.o google/protobuf/compiler/objectivec/.libs/objectivec_generator.o google/protobuf/compiler/objectivec/.libs/objectivec_helpers.o google/protobuf/compiler/objectivec/.libs/objectivec_map_field.o google/protobuf/compiler/objectivec/.libs/objectivec_message.o google/protobuf/compiler/objectivec/.libs/objectivec_message_field.o google/protobuf/compiler/objectivec/.libs/objectivec_oneof.o google/protobuf/compiler/objectivec/.libs/objectivec_primitive_field.o google/protobuf/compiler/php/.libs/php_generator.o google/protobuf/compiler/python/.libs/python_generator.o google/protobuf/compiler/ruby/.libs/ruby_generator.o google/protobuf/compiler/csharp/.libs/csharp_doc_comment.o google/protobuf/compiler/csharp/.libs/csharp_enum.o google/protobuf/compiler/csharp/.libs/csharp_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_field_base.o google/protobuf/compiler/csharp/.libs/csharp_generator.o google/protobuf/compiler/csharp/.libs/csharp_helpers.o google/protobuf/compiler/csharp/.libs/csharp_map_field.o google/protobuf/compiler/csharp/.libs/csharp_message.o google/protobuf/compiler/csharp/.libs/csharp_message_field.o google/protobuf/compiler/csharp/.libs/csharp_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_reflection_class.o google/protobuf/compiler/csharp/.libs/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_message_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_source_generator_base.o google/protobuf/compiler/csharp/.libs/csharp_wrapper_field.o   -Wl,-rpath -Wl,/work/src/github.com/google/protobuf/src/.libs -lpthread ./.libs/libprotobuf.so -L/usr/lib/gcc/x86_64-linux-gnu/4.8 -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../../lib -L/lib/x86_64-linux-gnu -L/lib/../lib -L/usr/lib/x86_64-linux-gnu -L/usr/lib/../lib -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-linux-gnu/4.8/crtendS.o /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crtn.o  -pthread -O2   -pthread -Wl,-soname -Wl,libprotoc.so.13 -o .libs/libprotoc.so.13.0.0
libtool: link: (cd ".libs" && rm -f "libprotoc.so.13" && ln -s "libprotoc.so.13.0.0" "libprotoc.so.13")
libtool: link: (cd ".libs" && rm -f "libprotoc.so" && ln -s "libprotoc.so.13.0.0" "libprotoc.so")
libtool: link: ar cru .libs/libprotoc.a  google/protobuf/compiler/code_generator.o google/protobuf/compiler/command_line_interface.o google/protobuf/compiler/plugin.o google/protobuf/compiler/plugin.pb.o google/protobuf/compiler/profile.pb.o google/protobuf/compiler/subprocess.o google/protobuf/compiler/zip_writer.o google/protobuf/compiler/cpp/cpp_enum.o google/protobuf/compiler/cpp/cpp_enum_field.o google/protobuf/compiler/cpp/cpp_extension.o google/protobuf/compiler/cpp/cpp_field.o google/protobuf/compiler/cpp/cpp_file.o google/protobuf/compiler/cpp/cpp_generator.o google/protobuf/compiler/cpp/cpp_helpers.o google/protobuf/compiler/cpp/cpp_map_field.o google/protobuf/compiler/cpp/cpp_message.o google/protobuf/compiler/cpp/cpp_message_field.o google/protobuf/compiler/cpp/cpp_primitive_field.o google/protobuf/compiler/cpp/cpp_service.o google/protobuf/compiler/cpp/cpp_string_field.o google/protobuf/compiler/java/java_context.o google/protobuf/compiler/java/java_enum.o google/protobuf/compiler/java/java_enum_lite.o google/protobuf/compiler/java/java_enum_field.o google/protobuf/compiler/java/java_enum_field_lite.o google/protobuf/compiler/java/java_extension.o google/protobuf/compiler/java/java_extension_lite.o google/protobuf/compiler/java/java_field.o google/protobuf/compiler/java/java_file.o google/protobuf/compiler/java/java_generator.o google/protobuf/compiler/java/java_generator_factory.o google/protobuf/compiler/java/java_helpers.o google/protobuf/compiler/java/java_lazy_message_field.o google/protobuf/compiler/java/java_lazy_message_field_lite.o google/protobuf/compiler/java/java_map_field.o google/protobuf/compiler/java/java_map_field_lite.o google/protobuf/compiler/java/java_message.o google/protobuf/compiler/java/java_message_lite.o google/protobuf/compiler/java/java_message_builder.o google/protobuf/compiler/java/java_message_builder_lite.o google/protobuf/compiler/java/java_message_field.o google/protobuf/compiler/java/java_message_field_lite.o google/protobuf/compiler/java/java_name_resolver.o google/protobuf/compiler/java/java_primitive_field.o google/protobuf/compiler/java/java_primitive_field_lite.o google/protobuf/compiler/java/java_shared_code_generator.o google/protobuf/compiler/java/java_service.o google/protobuf/compiler/java/java_string_field.o google/protobuf/compiler/java/java_string_field_lite.o google/protobuf/compiler/java/java_doc_comment.o google/protobuf/compiler/js/js_generator.o google/protobuf/compiler/js/well_known_types_embed.o google/protobuf/compiler/javanano/javanano_enum.o google/protobuf/compiler/javanano/javanano_enum_field.o google/protobuf/compiler/javanano/javanano_extension.o google/protobuf/compiler/javanano/javanano_field.o google/protobuf/compiler/javanano/javanano_file.o google/protobuf/compiler/javanano/javanano_generator.o google/protobuf/compiler/javanano/javanano_helpers.o google/protobuf/compiler/javanano/javanano_map_field.o google/protobuf/compiler/javanano/javanano_message.o google/protobuf/compiler/javanano/javanano_message_field.o google/protobuf/compiler/javanano/javanano_primitive_field.o google/protobuf/compiler/objectivec/objectivec_enum.o google/protobuf/compiler/objectivec/objectivec_enum_field.o google/protobuf/compiler/objectivec/objectivec_extension.o google/protobuf/compiler/objectivec/objectivec_field.o google/protobuf/compiler/objectivec/objectivec_file.o google/protobuf/compiler/objectivec/objectivec_generator.o google/protobuf/compiler/objectivec/objectivec_helpers.o google/protobuf/compiler/objectivec/objectivec_map_field.o google/protobuf/compiler/objectivec/objectivec_message.o google/protobuf/compiler/objectivec/objectivec_message_field.o google/protobuf/compiler/objectivec/objectivec_oneof.o google/protobuf/compiler/objectivec/objectivec_primitive_field.o google/protobuf/compiler/php/php_generator.o google/protobuf/compiler/python/python_generator.o google/protobuf/compiler/ruby/ruby_generator.o google/protobuf/compiler/csharp/csharp_doc_comment.o google/protobuf/compiler/csharp/csharp_enum.o google/protobuf/compiler/csharp/csharp_enum_field.o google/protobuf/compiler/csharp/csharp_field_base.o google/protobuf/compiler/csharp/csharp_generator.o google/protobuf/compiler/csharp/csharp_helpers.o google/protobuf/compiler/csharp/csharp_map_field.o google/protobuf/compiler/csharp/csharp_message.o google/protobuf/compiler/csharp/csharp_message_field.o google/protobuf/compiler/csharp/csharp_primitive_field.o google/protobuf/compiler/csharp/csharp_reflection_class.o google/protobuf/compiler/csharp/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/csharp_repeated_message_field.o google/protobuf/compiler/csharp/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/csharp_source_generator_base.o google/protobuf/compiler/csharp/csharp_wrapper_field.o
libtool: link: ranlib .libs/libprotoc.a
libtool: link: ( cd ".libs" && rm -f "libprotoc.la" && ln -s "../libprotoc.la" "libprotoc.la" )
/bin/sh ../libtool  --tag=CXX   --mode=link g++ -std=c++11 -pthread -DHAVE_PTHREAD=1  -Wall -Wno-sign-compare -O2 -g -DNDEBUG -pthread  -o protoc google/protobuf/compiler/main.o -lpthread libprotobuf.la libprotoc.la 
libtool: link: g++ -std=c++11 -pthread -DHAVE_PTHREAD=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -pthread -o .libs/protoc google/protobuf/compiler/main.o  -lpthread ./.libs/libprotobuf.so ./.libs/libprotoc.so -pthread
oldpwd=`pwd` && ( cd . && $oldpwd/protoc -I. --cpp_out=$oldpwd google/protobuf/any_test.proto google/protobuf/compiler/cpp/cpp_test_bad_identifiers.proto google/protobuf/map_lite_unittest.proto google/protobuf/map_proto2_unittest.proto google/protobuf/map_unittest.proto google/protobuf/unittest_arena.proto google/protobuf/unittest_custom_options.proto google/protobuf/unittest_drop_unknown_fields.proto google/protobuf/unittest_embed_optimize_for.proto google/protobuf/unittest_empty.proto google/protobuf/unittest_enormous_descriptor.proto google/protobuf/unittest_import_lite.proto google/protobuf/unittest_import.proto google/protobuf/unittest_import_public_lite.proto google/protobuf/unittest_import_public.proto google/protobuf/unittest_lazy_dependencies.proto google/protobuf/unittest_lazy_dependencies_custom_option.proto google/protobuf/unittest_lazy_dependencies_enum.proto google/protobuf/unittest_lite_imports_nonlite.proto google/protobuf/unittest_lite.proto google/protobuf/unittest_mset.proto google/protobuf/unittest_mset_wire_format.proto google/protobuf/unittest_no_arena_lite.proto google/protobuf/unittest_no_arena_import.proto google/protobuf/unittest_no_arena.proto google/protobuf/unittest_no_field_presence.proto google/protobuf/unittest_no_generic_services.proto google/protobuf/unittest_optimize_for.proto google/protobuf/unittest_preserve_unknown_enum2.proto google/protobuf/unittest_preserve_unknown_enum.proto google/protobuf/unittest.proto google/protobuf/unittest_proto3_arena.proto google/protobuf/unittest_proto3_arena_lite.proto google/protobuf/unittest_proto3_lite.proto google/protobuf/unittest_well_known_types.proto google/protobuf/util/internal/testdata/anys.proto google/protobuf/util/internal/testdata/books.proto google/protobuf/util/internal/testdata/default_value.proto google/protobuf/util/internal/testdata/default_value_test.proto google/protobuf/util/internal/testdata/field_mask.proto google/protobuf/util/internal/testdata/maps.proto google/protobuf/util/internal/testdata/oneofs.proto google/protobuf/util/internal/testdata/proto3.proto google/protobuf/util/internal/testdata/struct.proto google/protobuf/util/internal/testdata/timestamp_duration.proto google/protobuf/util/internal/testdata/wrappers.proto google/protobuf/util/json_format_proto3.proto google/protobuf/util/message_differencer_unittest.proto google/protobuf/compiler/cpp/cpp_test_large_enum_value.proto )
touch unittest_proto_middleman
make  all-am
make[3]: Entering directory `/work/src/github.com/google/protobuf/src'
/bin/sh ../libtool  --tag=CXX   --mode=link g++ -std=c++11 -pthread -DHAVE_PTHREAD=1  -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 13:0:0 -export-dynamic -no-undefined   -o libprotobuf-lite.la -rpath /usr/local/lib google/protobuf/stubs/atomicops_internals_x86_gcc.lo google/protobuf/stubs/atomicops_internals_x86_msvc.lo google/protobuf/stubs/bytestream.lo google/protobuf/stubs/common.lo google/protobuf/stubs/int128.lo google/protobuf/stubs/once.lo google/protobuf/stubs/status.lo google/protobuf/stubs/statusor.lo google/protobuf/stubs/stringpiece.lo google/protobuf/stubs/stringprintf.lo google/protobuf/stubs/structurally_valid.lo google/protobuf/stubs/strutil.lo google/protobuf/stubs/time.lo google/protobuf/arena.lo google/protobuf/arenastring.lo google/protobuf/extension_set.lo google/protobuf/generated_message_util.lo google/protobuf/message_lite.lo google/protobuf/repeated_field.lo google/protobuf/wire_format_lite.lo google/protobuf/io/coded_stream.lo google/protobuf/io/zero_copy_stream.lo google/protobuf/io/zero_copy_stream_impl_lite.lo -lpthread 
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crti.o /usr/lib/gcc/x86_64-linux-gnu/4.8/crtbeginS.o  google/protobuf/stubs/.libs/atomicops_internals_x86_gcc.o google/protobuf/stubs/.libs/atomicops_internals_x86_msvc.o google/protobuf/stubs/.libs/bytestream.o google/protobuf/stubs/.libs/common.o google/protobuf/stubs/.libs/int128.o google/protobuf/stubs/.libs/once.o google/protobuf/stubs/.libs/status.o google/protobuf/stubs/.libs/statusor.o google/protobuf/stubs/.libs/stringpiece.o google/protobuf/stubs/.libs/stringprintf.o google/protobuf/stubs/.libs/structurally_valid.o google/protobuf/stubs/.libs/strutil.o google/protobuf/stubs/.libs/time.o google/protobuf/.libs/arena.o google/protobuf/.libs/arenastring.o google/protobuf/.libs/extension_set.o google/protobuf/.libs/generated_message_util.o google/protobuf/.libs/message_lite.o google/protobuf/.libs/repeated_field.o google/protobuf/.libs/wire_format_lite.o google/protobuf/io/.libs/coded_stream.o google/protobuf/io/.libs/zero_copy_stream.o google/protobuf/io/.libs/zero_copy_stream_impl_lite.o   -lpthread -L/usr/lib/gcc/x86_64-linux-gnu/4.8 -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../../lib -L/lib/x86_64-linux-gnu -L/lib/../lib -L/usr/lib/x86_64-linux-gnu -L/usr/lib/../lib -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-linux-gnu/4.8/crtendS.o /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crtn.o  -pthread -O2   -pthread -Wl,-soname -Wl,libprotobuf-lite.so.13 -o .libs/libprotobuf-lite.so.13.0.0
libtool: link: (cd ".libs" && rm -f "libprotobuf-lite.so.13" && ln -s "libprotobuf-lite.so.13.0.0" "libprotobuf-lite.so.13")
libtool: link: (cd ".libs" && rm -f "libprotobuf-lite.so" && ln -s "libprotobuf-lite.so.13.0.0" "libprotobuf-lite.so")
libtool: link: ar cru .libs/libprotobuf-lite.a  google/protobuf/stubs/atomicops_internals_x86_gcc.o google/protobuf/stubs/atomicops_internals_x86_msvc.o google/protobuf/stubs/bytestream.o google/protobuf/stubs/common.o google/protobuf/stubs/int128.o google/protobuf/stubs/once.o google/protobuf/stubs/status.o google/protobuf/stubs/statusor.o google/protobuf/stubs/stringpiece.o google/protobuf/stubs/stringprintf.o google/protobuf/stubs/structurally_valid.o google/protobuf/stubs/strutil.o google/protobuf/stubs/time.o google/protobuf/arena.o google/protobuf/arenastring.o google/protobuf/extension_set.o google/protobuf/generated_message_util.o google/protobuf/message_lite.o google/protobuf/repeated_field.o google/protobuf/wire_format_lite.o google/protobuf/io/coded_stream.o google/protobuf/io/zero_copy_stream.o google/protobuf/io/zero_copy_stream_impl_lite.o
libtool: link: ranlib .libs/libprotobuf-lite.a
libtool: link: ( cd ".libs" && rm -f "libprotobuf-lite.la" && ln -s "../libprotobuf-lite.la" "libprotobuf-lite.la" )
make[3]: Leaving directory `/work/src/github.com/google/protobuf/src'
make[2]: Leaving directory `/work/src/github.com/google/protobuf/src'
make[1]: Leaving directory `/work/src/github.com/google/protobuf'
```

Install
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ sudo make install
Making install in .
make[1]: Entering directory `/work/src/github.com/google/protobuf'
make[2]: Entering directory `/work/src/github.com/google/protobuf'
make[2]: Nothing to be done for `install-exec-am'.
 /bin/mkdir -p '/usr/local/lib/pkgconfig'
 /usr/bin/install -c -m 644 protobuf.pc protobuf-lite.pc '/usr/local/lib/pkgconfig'
make[2]: Leaving directory `/work/src/github.com/google/protobuf'
make[1]: Leaving directory `/work/src/github.com/google/protobuf'
Making install in src
make[1]: Entering directory `/work/src/github.com/google/protobuf/src'
make  install-am
make[2]: Entering directory `/work/src/github.com/google/protobuf/src'
make[3]: Entering directory `/work/src/github.com/google/protobuf/src'
 /bin/mkdir -p '/usr/local/lib'
 /bin/sh ../libtool   --mode=install /usr/bin/install -c   libprotobuf-lite.la libprotobuf.la libprotoc.la '/usr/local/lib'
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.so.13.0.0 /usr/local/lib/libprotobuf-lite.so.13.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf-lite.so.13.0.0 libprotobuf-lite.so.13 || { rm -f libprotobuf-lite.so.13 && ln -s libprotobuf-lite.so.13.0.0 libprotobuf-lite.so.13; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf-lite.so.13.0.0 libprotobuf-lite.so || { rm -f libprotobuf-lite.so && ln -s libprotobuf-lite.so.13.0.0 libprotobuf-lite.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.lai /usr/local/lib/libprotobuf-lite.la
libtool: install: /usr/bin/install -c .libs/libprotobuf.so.13.0.0 /usr/local/lib/libprotobuf.so.13.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf.so.13.0.0 libprotobuf.so.13 || { rm -f libprotobuf.so.13 && ln -s libprotobuf.so.13.0.0 libprotobuf.so.13; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf.so.13.0.0 libprotobuf.so || { rm -f libprotobuf.so && ln -s libprotobuf.so.13.0.0 libprotobuf.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotobuf.lai /usr/local/lib/libprotobuf.la
libtool: install: warning: relinking `libprotoc.la'
libtool: install: (cd /work/src/github.com/google/protobuf/src; /bin/sh /work/src/github.com/google/protobuf/libtool  --tag CXX --mode=relink g++ -std=c++11 -pthread -DHAVE_PTHREAD=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 13:0:0 -export-dynamic -no-undefined -o libprotoc.la -rpath /usr/local/lib google/protobuf/compiler/code_generator.lo google/protobuf/compiler/command_line_interface.lo google/protobuf/compiler/plugin.lo google/protobuf/compiler/plugin.pb.lo google/protobuf/compiler/profile.pb.lo google/protobuf/compiler/subprocess.lo google/protobuf/compiler/zip_writer.lo google/protobuf/compiler/cpp/cpp_enum.lo google/protobuf/compiler/cpp/cpp_enum_field.lo google/protobuf/compiler/cpp/cpp_extension.lo google/protobuf/compiler/cpp/cpp_field.lo google/protobuf/compiler/cpp/cpp_file.lo google/protobuf/compiler/cpp/cpp_generator.lo google/protobuf/compiler/cpp/cpp_helpers.lo google/protobuf/compiler/cpp/cpp_map_field.lo google/protobuf/compiler/cpp/cpp_message.lo google/protobuf/compiler/cpp/cpp_message_field.lo google/protobuf/compiler/cpp/cpp_primitive_field.lo google/protobuf/compiler/cpp/cpp_service.lo google/protobuf/compiler/cpp/cpp_string_field.lo google/protobuf/compiler/java/java_context.lo google/protobuf/compiler/java/java_enum.lo google/protobuf/compiler/java/java_enum_lite.lo google/protobuf/compiler/java/java_enum_field.lo google/protobuf/compiler/java/java_enum_field_lite.lo google/protobuf/compiler/java/java_extension.lo google/protobuf/compiler/java/java_extension_lite.lo google/protobuf/compiler/java/java_field.lo google/protobuf/compiler/java/java_file.lo google/protobuf/compiler/java/java_generator.lo google/protobuf/compiler/java/java_generator_factory.lo google/protobuf/compiler/java/java_helpers.lo google/protobuf/compiler/java/java_lazy_message_field.lo google/protobuf/compiler/java/java_lazy_message_field_lite.lo google/protobuf/compiler/java/java_map_field.lo google/protobuf/compiler/java/java_map_field_lite.lo google/protobuf/compiler/java/java_message.lo google/protobuf/compiler/java/java_message_lite.lo google/protobuf/compiler/java/java_message_builder.lo google/protobuf/compiler/java/java_message_builder_lite.lo google/protobuf/compiler/java/java_message_field.lo google/protobuf/compiler/java/java_message_field_lite.lo google/protobuf/compiler/java/java_name_resolver.lo google/protobuf/compiler/java/java_primitive_field.lo google/protobuf/compiler/java/java_primitive_field_lite.lo google/protobuf/compiler/java/java_shared_code_generator.lo google/protobuf/compiler/java/java_service.lo google/protobuf/compiler/java/java_string_field.lo google/protobuf/compiler/java/java_string_field_lite.lo google/protobuf/compiler/java/java_doc_comment.lo google/protobuf/compiler/js/js_generator.lo google/protobuf/compiler/js/well_known_types_embed.lo google/protobuf/compiler/javanano/javanano_enum.lo google/protobuf/compiler/javanano/javanano_enum_field.lo google/protobuf/compiler/javanano/javanano_extension.lo google/protobuf/compiler/javanano/javanano_field.lo google/protobuf/compiler/javanano/javanano_file.lo google/protobuf/compiler/javanano/javanano_generator.lo google/protobuf/compiler/javanano/javanano_helpers.lo google/protobuf/compiler/javanano/javanano_map_field.lo google/protobuf/compiler/javanano/javanano_message.lo google/protobuf/compiler/javanano/javanano_message_field.lo google/protobuf/compiler/javanano/javanano_primitive_field.lo google/protobuf/compiler/objectivec/objectivec_enum.lo google/protobuf/compiler/objectivec/objectivec_enum_field.lo google/protobuf/compiler/objectivec/objectivec_extension.lo google/protobuf/compiler/objectivec/objectivec_field.lo google/protobuf/compiler/objectivec/objectivec_file.lo google/protobuf/compiler/objectivec/objectivec_generator.lo google/protobuf/compiler/objectivec/objectivec_helpers.lo google/protobuf/compiler/objectivec/objectivec_map_field.lo google/protobuf/compiler/objectivec/objectivec_message.lo google/protobuf/compiler/objectivec/objectivec_message_field.lo google/protobuf/compiler/objectivec/objectivec_oneof.lo google/protobuf/compiler/objectivec/objectivec_primitive_field.lo google/protobuf/compiler/php/php_generator.lo google/protobuf/compiler/python/python_generator.lo google/protobuf/compiler/ruby/ruby_generator.lo google/protobuf/compiler/csharp/csharp_doc_comment.lo google/protobuf/compiler/csharp/csharp_enum.lo google/protobuf/compiler/csharp/csharp_enum_field.lo google/protobuf/compiler/csharp/csharp_field_base.lo google/protobuf/compiler/csharp/csharp_generator.lo google/protobuf/compiler/csharp/csharp_helpers.lo google/protobuf/compiler/csharp/csharp_map_field.lo google/protobuf/compiler/csharp/csharp_message.lo google/protobuf/compiler/csharp/csharp_message_field.lo google/protobuf/compiler/csharp/csharp_primitive_field.lo google/protobuf/compiler/csharp/csharp_reflection_class.lo google/protobuf/compiler/csharp/csharp_repeated_enum_field.lo google/protobuf/compiler/csharp/csharp_repeated_message_field.lo google/protobuf/compiler/csharp/csharp_repeated_primitive_field.lo google/protobuf/compiler/csharp/csharp_source_generator_base.lo google/protobuf/compiler/csharp/csharp_wrapper_field.lo -lpthread libprotobuf.la )
libtool: relink: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crti.o /usr/lib/gcc/x86_64-linux-gnu/4.8/crtbeginS.o  google/protobuf/compiler/.libs/code_generator.o google/protobuf/compiler/.libs/command_line_interface.o google/protobuf/compiler/.libs/plugin.o google/protobuf/compiler/.libs/plugin.pb.o google/protobuf/compiler/.libs/profile.pb.o google/protobuf/compiler/.libs/subprocess.o google/protobuf/compiler/.libs/zip_writer.o google/protobuf/compiler/cpp/.libs/cpp_enum.o google/protobuf/compiler/cpp/.libs/cpp_enum_field.o google/protobuf/compiler/cpp/.libs/cpp_extension.o google/protobuf/compiler/cpp/.libs/cpp_field.o google/protobuf/compiler/cpp/.libs/cpp_file.o google/protobuf/compiler/cpp/.libs/cpp_generator.o google/protobuf/compiler/cpp/.libs/cpp_helpers.o google/protobuf/compiler/cpp/.libs/cpp_map_field.o google/protobuf/compiler/cpp/.libs/cpp_message.o google/protobuf/compiler/cpp/.libs/cpp_message_field.o google/protobuf/compiler/cpp/.libs/cpp_primitive_field.o google/protobuf/compiler/cpp/.libs/cpp_service.o google/protobuf/compiler/cpp/.libs/cpp_string_field.o google/protobuf/compiler/java/.libs/java_context.o google/protobuf/compiler/java/.libs/java_enum.o google/protobuf/compiler/java/.libs/java_enum_lite.o google/protobuf/compiler/java/.libs/java_enum_field.o google/protobuf/compiler/java/.libs/java_enum_field_lite.o google/protobuf/compiler/java/.libs/java_extension.o google/protobuf/compiler/java/.libs/java_extension_lite.o google/protobuf/compiler/java/.libs/java_field.o google/protobuf/compiler/java/.libs/java_file.o google/protobuf/compiler/java/.libs/java_generator.o google/protobuf/compiler/java/.libs/java_generator_factory.o google/protobuf/compiler/java/.libs/java_helpers.o google/protobuf/compiler/java/.libs/java_lazy_message_field.o google/protobuf/compiler/java/.libs/java_lazy_message_field_lite.o google/protobuf/compiler/java/.libs/java_map_field.o google/protobuf/compiler/java/.libs/java_map_field_lite.o google/protobuf/compiler/java/.libs/java_message.o google/protobuf/compiler/java/.libs/java_message_lite.o google/protobuf/compiler/java/.libs/java_message_builder.o google/protobuf/compiler/java/.libs/java_message_builder_lite.o google/protobuf/compiler/java/.libs/java_message_field.o google/protobuf/compiler/java/.libs/java_message_field_lite.o google/protobuf/compiler/java/.libs/java_name_resolver.o google/protobuf/compiler/java/.libs/java_primitive_field.o google/protobuf/compiler/java/.libs/java_primitive_field_lite.o google/protobuf/compiler/java/.libs/java_shared_code_generator.o google/protobuf/compiler/java/.libs/java_service.o google/protobuf/compiler/java/.libs/java_string_field.o google/protobuf/compiler/java/.libs/java_string_field_lite.o google/protobuf/compiler/java/.libs/java_doc_comment.o google/protobuf/compiler/js/.libs/js_generator.o google/protobuf/compiler/js/.libs/well_known_types_embed.o google/protobuf/compiler/javanano/.libs/javanano_enum.o google/protobuf/compiler/javanano/.libs/javanano_enum_field.o google/protobuf/compiler/javanano/.libs/javanano_extension.o google/protobuf/compiler/javanano/.libs/javanano_field.o google/protobuf/compiler/javanano/.libs/javanano_file.o google/protobuf/compiler/javanano/.libs/javanano_generator.o google/protobuf/compiler/javanano/.libs/javanano_helpers.o google/protobuf/compiler/javanano/.libs/javanano_map_field.o google/protobuf/compiler/javanano/.libs/javanano_message.o google/protobuf/compiler/javanano/.libs/javanano_message_field.o google/protobuf/compiler/javanano/.libs/javanano_primitive_field.o google/protobuf/compiler/objectivec/.libs/objectivec_enum.o google/protobuf/compiler/objectivec/.libs/objectivec_enum_field.o google/protobuf/compiler/objectivec/.libs/objectivec_extension.o google/protobuf/compiler/objectivec/.libs/objectivec_field.o google/protobuf/compiler/objectivec/.libs/objectivec_file.o google/protobuf/compiler/objectivec/.libs/objectivec_generator.o google/protobuf/compiler/objectivec/.libs/objectivec_helpers.o google/protobuf/compiler/objectivec/.libs/objectivec_map_field.o google/protobuf/compiler/objectivec/.libs/objectivec_message.o google/protobuf/compiler/objectivec/.libs/objectivec_message_field.o google/protobuf/compiler/objectivec/.libs/objectivec_oneof.o google/protobuf/compiler/objectivec/.libs/objectivec_primitive_field.o google/protobuf/compiler/php/.libs/php_generator.o google/protobuf/compiler/python/.libs/python_generator.o google/protobuf/compiler/ruby/.libs/ruby_generator.o google/protobuf/compiler/csharp/.libs/csharp_doc_comment.o google/protobuf/compiler/csharp/.libs/csharp_enum.o google/protobuf/compiler/csharp/.libs/csharp_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_field_base.o google/protobuf/compiler/csharp/.libs/csharp_generator.o google/protobuf/compiler/csharp/.libs/csharp_helpers.o google/protobuf/compiler/csharp/.libs/csharp_map_field.o google/protobuf/compiler/csharp/.libs/csharp_message.o google/protobuf/compiler/csharp/.libs/csharp_message_field.o google/protobuf/compiler/csharp/.libs/csharp_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_reflection_class.o google/protobuf/compiler/csharp/.libs/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_message_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_source_generator_base.o google/protobuf/compiler/csharp/.libs/csharp_wrapper_field.o   -lpthread -L/usr/local/lib -lprotobuf -L/usr/lib/gcc/x86_64-linux-gnu/4.8 -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../../../lib -L/lib/x86_64-linux-gnu -L/lib/../lib -L/usr/lib/x86_64-linux-gnu -L/usr/lib/../lib -L/usr/lib/gcc/x86_64-linux-gnu/4.8/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-linux-gnu/4.8/crtendS.o /usr/lib/gcc/x86_64-linux-gnu/4.8/../../../x86_64-linux-gnu/crtn.o  -pthread -O2   -pthread -Wl,-soname -Wl,libprotoc.so.13 -o .libs/libprotoc.so.13.0.0
libtool: install: /usr/bin/install -c .libs/libprotoc.so.13.0.0T /usr/local/lib/libprotoc.so.13.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotoc.so.13.0.0 libprotoc.so.13 || { rm -f libprotoc.so.13 && ln -s libprotoc.so.13.0.0 libprotoc.so.13; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotoc.so.13.0.0 libprotoc.so || { rm -f libprotoc.so && ln -s libprotoc.so.13.0.0 libprotoc.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotoc.lai /usr/local/lib/libprotoc.la
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.a /usr/local/lib/libprotobuf-lite.a
libtool: install: chmod 644 /usr/local/lib/libprotobuf-lite.a
libtool: install: ranlib /usr/local/lib/libprotobuf-lite.a
libtool: install: /usr/bin/install -c .libs/libprotobuf.a /usr/local/lib/libprotobuf.a
libtool: install: chmod 644 /usr/local/lib/libprotobuf.a
libtool: install: ranlib /usr/local/lib/libprotobuf.a
libtool: install: /usr/bin/install -c .libs/libprotoc.a /usr/local/lib/libprotoc.a
libtool: install: chmod 644 /usr/local/lib/libprotoc.a
libtool: install: ranlib /usr/local/lib/libprotoc.a
libtool: finish: PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin:/sbin" ldconfig -n /usr/local/lib
----------------------------------------------------------------------
Libraries have been installed in:
   /usr/local/lib

If you ever happen to want to link against installed libraries
in a given directory, LIBDIR, you must either use libtool, and
specify the full pathname of the library, or use the `-LLIBDIR'
flag during linking and do at least one of the following:
   - add LIBDIR to the `LD_LIBRARY_PATH' environment variable
     during execution
   - add LIBDIR to the `LD_RUN_PATH' environment variable
     during linking
   - use the `-Wl,-rpath -Wl,LIBDIR' linker flag
   - have your system administrator add LIBDIR to `/etc/ld.so.conf'

See any operating system documentation about shared libraries for
more information, such as the ld(1) and ld.so(8) manual pages.
----------------------------------------------------------------------
 /bin/mkdir -p '/usr/local/bin'
  /bin/sh ../libtool   --mode=install /usr/bin/install -c protoc '/usr/local/bin'
libtool: install: /usr/bin/install -c .libs/protoc /usr/local/bin/protoc
 /bin/mkdir -p '/usr/local/include'
 /bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/descriptor.proto google/protobuf/any.proto google/protobuf/api.proto google/protobuf/duration.proto google/protobuf/empty.proto google/protobuf/field_mask.proto google/protobuf/source_context.proto google/protobuf/struct.proto google/protobuf/timestamp.proto google/protobuf/type.proto google/protobuf/wrappers.proto '/usr/local/include/google/protobuf'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler'
 /usr/bin/install -c -m 644  google/protobuf/compiler/plugin.proto google/protobuf/compiler/profile.proto '/usr/local/include/google/protobuf/compiler'
 /bin/mkdir -p '/usr/local/include'
 /bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/any.pb.h google/protobuf/api.pb.h google/protobuf/any.h google/protobuf/arena.h google/protobuf/arenastring.h google/protobuf/descriptor_database.h google/protobuf/descriptor.h google/protobuf/descriptor.pb.h google/protobuf/duration.pb.h google/protobuf/dynamic_message.h google/protobuf/empty.pb.h google/protobuf/extension_set.h google/protobuf/field_mask.pb.h google/protobuf/generated_enum_reflection.h google/protobuf/generated_enum_util.h google/protobuf/generated_message_reflection.h google/protobuf/generated_message_table_driven.h google/protobuf/generated_message_util.h google/protobuf/has_bits.h google/protobuf/map_entry.h google/protobuf/map_entry_lite.h google/protobuf/map_field.h google/protobuf/map_field_inl.h google/protobuf/map_field_lite.h google/protobuf/map.h google/protobuf/map_type_handler.h google/protobuf/message.h google/protobuf/message_lite.h google/protobuf/metadata.h google/protobuf/metadata_lite.h google/protobuf/reflection.h google/protobuf/reflection_ops.h google/protobuf/repeated_field.h google/protobuf/service.h google/protobuf/source_context.pb.h google/protobuf/struct.pb.h google/protobuf/text_format.h google/protobuf/timestamp.pb.h google/protobuf/type.pb.h google/protobuf/unknown_field_set.h '/usr/local/include/google/protobuf'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/javanano'
 /usr/bin/install -c -m 644  google/protobuf/compiler/javanano/javanano_generator.h '/usr/local/include/google/protobuf/compiler/javanano'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/java'
 /usr/bin/install -c -m 644  google/protobuf/compiler/java/java_generator.h google/protobuf/compiler/java/java_names.h '/usr/local/include/google/protobuf/compiler/java'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/cpp'
 /usr/bin/install -c -m 644  google/protobuf/compiler/cpp/cpp_generator.h '/usr/local/include/google/protobuf/compiler/cpp'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/python'
 /usr/bin/install -c -m 644  google/protobuf/compiler/python/python_generator.h '/usr/local/include/google/protobuf/compiler/python'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/js'
 /usr/bin/install -c -m 644  google/protobuf/compiler/js/js_generator.h google/protobuf/compiler/js/well_known_types_embed.h '/usr/local/include/google/protobuf/compiler/js'
 /bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/wire_format.h google/protobuf/wire_format_lite.h google/protobuf/wire_format_lite_inl.h google/protobuf/wrappers.pb.h '/usr/local/include/google/protobuf'
 /bin/mkdir -p '/usr/local/include/google/protobuf/stubs'
 /usr/bin/install -c -m 644  google/protobuf/stubs/atomic_sequence_num.h google/protobuf/stubs/atomicops.h google/protobuf/stubs/atomicops_internals_power.h google/protobuf/stubs/atomicops_internals_ppc_gcc.h google/protobuf/stubs/atomicops_internals_arm64_gcc.h google/protobuf/stubs/atomicops_internals_arm_gcc.h google/protobuf/stubs/atomicops_internals_arm_qnx.h google/protobuf/stubs/atomicops_internals_atomicword_compat.h google/protobuf/stubs/atomicops_internals_generic_c11_atomic.h google/protobuf/stubs/atomicops_internals_generic_gcc.h google/protobuf/stubs/atomicops_internals_mips_gcc.h google/protobuf/stubs/atomicops_internals_solaris.h google/protobuf/stubs/atomicops_internals_tsan.h google/protobuf/stubs/atomicops_internals_x86_gcc.h google/protobuf/stubs/atomicops_internals_x86_msvc.h google/protobuf/stubs/callback.h google/protobuf/stubs/bytestream.h google/protobuf/stubs/casts.h google/protobuf/stubs/common.h google/protobuf/stubs/fastmem.h google/protobuf/stubs/hash.h google/protobuf/stubs/logging.h google/protobuf/stubs/macros.h google/protobuf/stubs/mutex.h google/protobuf/stubs/once.h google/protobuf/stubs/platform_macros.h google/protobuf/stubs/port.h google/protobuf/stubs/scoped_ptr.h google/protobuf/stubs/shared_ptr.h google/protobuf/stubs/singleton.h google/protobuf/stubs/status.h google/protobuf/stubs/stl_util.h google/protobuf/stubs/stringpiece.h google/protobuf/stubs/template_util.h google/protobuf/stubs/type_traits.h '/usr/local/include/google/protobuf/stubs'
 /bin/mkdir -p '/usr/local/include/google/protobuf/util'
 /usr/bin/install -c -m 644  google/protobuf/util/type_resolver.h google/protobuf/util/delimited_message_util.h google/protobuf/util/field_comparator.h google/protobuf/util/field_mask_util.h google/protobuf/util/json_util.h google/protobuf/util/time_util.h google/protobuf/util/type_resolver_util.h google/protobuf/util/message_differencer.h '/usr/local/include/google/protobuf/util'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/php'
 /usr/bin/install -c -m 644  google/protobuf/compiler/php/php_generator.h '/usr/local/include/google/protobuf/compiler/php'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler'
 /usr/bin/install -c -m 644  google/protobuf/compiler/code_generator.h google/protobuf/compiler/command_line_interface.h google/protobuf/compiler/importer.h google/protobuf/compiler/parser.h google/protobuf/compiler/plugin.h google/protobuf/compiler/plugin.pb.h google/protobuf/compiler/profile.pb.h '/usr/local/include/google/protobuf/compiler'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/ruby'
 /usr/bin/install -c -m 644  google/protobuf/compiler/ruby/ruby_generator.h '/usr/local/include/google/protobuf/compiler/ruby'
 /bin/mkdir -p '/usr/local/include/google/protobuf/io'
 /usr/bin/install -c -m 644  google/protobuf/io/coded_stream.h google/protobuf/io/printer.h google/protobuf/io/strtod.h google/protobuf/io/tokenizer.h google/protobuf/io/zero_copy_stream.h google/protobuf/io/zero_copy_stream_impl.h google/protobuf/io/zero_copy_stream_impl_lite.h '/usr/local/include/google/protobuf/io'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/csharp'
 /usr/bin/install -c -m 644  google/protobuf/compiler/csharp/csharp_generator.h google/protobuf/compiler/csharp/csharp_names.h '/usr/local/include/google/protobuf/compiler/csharp'
 /bin/mkdir -p '/usr/local/include/google/protobuf/compiler/objectivec'
 /usr/bin/install -c -m 644  google/protobuf/compiler/objectivec/objectivec_generator.h google/protobuf/compiler/objectivec/objectivec_helpers.h '/usr/local/include/google/protobuf/compiler/objectivec'
make[3]: Leaving directory `/work/src/github.com/google/protobuf/src'
make[2]: Leaving directory `/work/src/github.com/google/protobuf/src'
make[1]: Leaving directory `/work/src/github.com/google/protobuf/src'
```

Refesh _shared libraries_
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ sudo ldconfig
```

Check
```
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ which protoc
/usr/local/bin/protoc
vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/google/protobuf$ protoc --version
libprotoc 3.3.0
```