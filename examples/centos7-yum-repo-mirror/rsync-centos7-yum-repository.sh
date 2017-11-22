#!/bin/bash -e

# https://www.centos.org/download/mirrors/

# ALIYUN_MIRROR_URL="http://mirrors.aliyun.com/centos/"
CENTOS_MIRROR_URL="http://mirror.centos.org/centos/"
# YUN_IDC_MIRROR="rsync://mirrors.yun-idc.com/centos/"
JAPAN_JAIST="rsync://ftp.jaist.ac.jp/pub/Linux/CentOS/"

HUAZHONG_UNIVERSITY="rsync://mirrors.hust.edu.cn/centos/"
TSINGHUA_UNIVERSITY="rsync://mirrors.tuna.tsinghua.edu.cn/centos/"

# rsync -av --list-only --delete-excluded rsync://mirrors.tuna.tsinghua.edu.cn/centos/7

# SRC_URL="${JAPAN_JAIST}7"
SRC_URL="${TSINGHUA_UNIVERSITY}7"

SRC_NAME=${SRC_URL//\//\%2F}
SRC_NAME=${SRC_NAME/\:/\%3A}

REPO_PATH='rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7'
DEST_HOME=$(dirname "${BASH_SOURCE[0]}")

# cd $DEST_HOME && ln -sf $DEST_HOME/$REPO_PATH $DEST_HOME/$SRC_NAME 
if [[ ! -d $DEST_HOME/$REPO_PATH ]]; then mkdir -p $DEST_HOME/$REPO_PATH; fi
  
REPO_LIST='centosplus cloud cr extras fasttrack paas sclo storage updates virt'
for REPO_ITEM in $REPO_LIST; do 
    rsync -avhHSP --delete --delete-excluded \
        --filter=". ${DEST_HOME}/centos-rsync.filter" \
        --filter=":n- .rsync.filter" \
        --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
        --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
        ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH ;
done		
 
while [ $# -gt 0 ]; do
	case "$1" in
		os)
REPO_ITEM='os'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		isos)
REPO_ITEM='isos'
rsync -avh --delete \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		atomic)
REPO_ITEM='atomic'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		configmanagement)
REPO_ITEM='configmanagement'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		dotnet)
REPO_ITEM='dotnet'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		opstools)
REPO_ITEM='opstools'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
		rt)
REPO_ITEM='rt'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$REPO_PATH
			;;
        *)
		    echo "unrecognized repo item: $1"
		    ;;
	esac
	shift
done
