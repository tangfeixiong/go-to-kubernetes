#!/bin/bash -e

# https://www.centos.org/download/mirrors/

ALIYUN_MIRROR_URL="http://mirrors.aliyun.com/centos/7"
CENTOS_MIRROR_URL="http://mirror.centos.org/centos/7"
YUN_IDC_URL="rsync://mirrors.yun-idc.com/centos/7"

SRC_URL="rsync://mirrors.yun-idc.com/centos/7"

TGT=${SRC_URL//\//\%2F}
TGT=${TGT/\:/\%3A}

DEST_HOME=$(dirname "${BASH_SOURCE[0]}")

if [[ ! -d $DEST_HOME/$TGT ]]; then mkdir -p $DEST_HOME/$TGT; fi

#rsync -av --list-only --delete-excluded rsync://mirrors.yun-idc.com/centos/7
  
REPO_LIST='centosplus cloud cr extras fasttrack os paas sclo storage updates virt'
for REPO_ITEM in $REPO_LIST; do 
    rsync -avhHSP --delete --delete-excluded \
        --filter=". ${DEST_HOME}/centos-rsync.filter" \
        --filter=":n- .rsync.filter" \
        --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
        --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
        ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$TGT ;
done		
  
REPO_ITEM='os'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$TGT  
  
REPO_ITEM='isos'
rsync -avh --delete \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$TGT
  
REPO_ITEM='atomic'
rsync -avh --delete --delete-excluded \
  --filter=". ${DEST_HOME}/centos-rsync.filter" \
  --filter=":n- .rsync.filter" \
  --exclude-from="${DEST_HOME}/${REPO_ITEM}.exclude" \
  --include-from="${DEST_HOME}/${REPO_ITEM}.include" \
  ${SRC_URL}/${REPO_ITEM} ${DEST_HOME}/$TGT
