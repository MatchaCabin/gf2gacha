<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {GetGameInfo, GetPoolInfo, GetUserList, IncrementalUpdatePoolInfo, MergeEreRecord} from "../wailsjs/go/main/App";
import PoolCard from "./components/PoolCard.vue";
import {model} from "../wailsjs/go/models";
import 'element-plus/es/components/message/style/css'
import {ElMessage} from "element-plus";
import {Connection, CopyDocument} from "@element-plus/icons-vue";
import {ClipboardSetText} from "../wailsjs/runtime";
import Pool = model.Pool;
import Info = model.Info;

const currentUid = ref("");
const uidList = ref<string[]>([]);
const poolList = ref<Pool[]>([]);
const gameInfo = ref<Info>({})
const loading = ref(false);
const dialogInfoVisible = ref(false)
const dialogEreVisible = ref(false)


const getUidList = async () => {
  await GetUserList().then(result => {
    if (result) {
      uidList.value = result
    }
  })
}

const getPoolInfo = async (poolType: number) => {
  await GetPoolInfo(currentUid.value, poolType).then(result => {
    let list = poolList.value
    list.push(result)
    poolList.value = list
  })
}

const getAllPoolInfo = async () => {
  poolList.value = []
  await getPoolInfo(3)
  await getPoolInfo(4)
  await getPoolInfo(1)
  await getPoolInfo(5)
  await getPoolInfo(8)
}

const incrementalUpdatePoolInfo = async () => {
  loading.value = true
  await IncrementalUpdatePoolInfo().then(result => {
    if (result != "") {
      if (!uidList.value.includes(result)) {
        uidList.value.push(result)
      }
      currentUid.value = result
    }
  }).catch(err => {
    console.log(err)
  })
  await getAllPoolInfo()
  ElMessage({
    message: '更新成功',
    type: 'success',
    plain: true,
    showClose: true,
    duration: 1000
  })
  loading.value = false
}

const openInfoDialog = async () => {
  await GetGameInfo().then(result => {
    gameInfo.value = result
  })
  dialogInfoVisible.value = true
}

const mergeEreRecord = async () => {
  loading.value = true
  await MergeEreRecord(currentUid.value).then(() => {
    ElMessage({message: '合并成功', type: 'success', plain: true, showClose: true, duration: 1000})
  }).catch(() => {
    ElMessage({message: '合并发生错误', type: 'error', plain: true, showClose: true, duration: 1000})
  })
  await getAllPoolInfo()
  loading.value = false
}

const copyUid = () => {
  ClipboardSetText(gameInfo.value.Uid)
  ElMessage({message: 'UID已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

const copyGachaUrl = () => {
  ClipboardSetText(gameInfo.value.GachaUrl)
  ElMessage({message: '抽卡链接已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

const copyAccessToken = () => {
  ClipboardSetText(gameInfo.value.AccessToken)
  ElMessage({message: 'AccessToken已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

onMounted(async () => {
  await getUidList()
  if (uidList.value.length > 0) {
    currentUid.value = uidList.value[0]
    await getAllPoolInfo()
  }
})

</script>

<template>
  <div class="h-dvh w-full flex flex-col p-4 gap-4" v-loading="loading" element-loading-text="Loading...">
    <div class="flex">
      <div class="grow">
        <el-button type="success" class="font-bold" @click="incrementalUpdatePoolInfo">增量更新</el-button>
        <el-button type="primary" class="font-bold" disabled>全量更新</el-button>
        <el-popconfirm class="text-red-600" width="360" confirm-button-text="确定" cancel-button-text="取消" :title="`确定将数据合并进当前用户(UID:${currentUid})?`" @confirm="mergeEreRecord">
          <template #reference>
            <el-button type="danger" class="font-bold" :disabled="!currentUid">导入ERE数据</el-button>
          </template>
        </el-popconfirm>
      </div>
      <div class="flex items-center gap-2">
        <div>UID:</div>
        <el-select v-model="currentUid" class="w-28" @change="getAllPoolInfo">
          <el-option v-for="uid in uidList" :key="uid" :label="uid" :value="uid"/>
        </el-select>
        <el-button text :icon="Connection" circle @click="openInfoDialog"/>
      </div>
    </div>
    <div class="w-full flex flex-wrap justify-between">
      <PoolCard v-for="pool in poolList" :pool="pool"></PoolCard>
    </div>
    <el-dialog v-model="dialogInfoVisible" width="600">
      <template #title>
        <div class="text-xl font-bold">当前日志信息</div>
      </template>
      <div class="flex flex-col gap-4">
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">UID</div>
          <el-input class="grow" readonly v-model="gameInfo.Uid"/>
          <el-button text :icon="CopyDocument" circle @click="copyUid"/>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">抽卡链接</div>
          <el-input class="grow" readonly v-model="gameInfo.GachaUrl"/>
          <el-button text :icon="CopyDocument" circle @click="copyGachaUrl"/>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">AccessToken</div>
          <el-input class="grow" readonly type="password" v-model="gameInfo.AccessToken"/>
          <el-button text :icon="CopyDocument" circle @click="copyAccessToken"/>
        </div>
        <el-alert title="请勿随意泄露AccessToken，虽然操作上有难度，但理论上可以融号" type="warning" show-icon :closable="false"></el-alert>
      </div>
    </el-dialog>
  </div>
</template>

<style>

</style>
