<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {ApplyUpdate, CheckUpdate, GetLogInfo, GetPoolInfo, GetUserList, HandleCommunityTasks, IncrementalUpdatePoolInfo, MergeEreRecord} from "../wailsjs/go/main/App";
import PoolCard from "./components/PoolCard.vue";
import {model} from "../wailsjs/go/models";
import 'element-plus/es/components/message/style/css'
import {ElMessage} from "element-plus";
import {Connection, CopyDocument} from "@element-plus/icons-vue";
import {ClipboardSetText} from "../wailsjs/runtime";
import Pool = model.Pool;
import LogInfo = model.LogInfo;

const currentUid = ref("");
const uidList = ref<string[]>([]);
const poolList = ref<Pool[]>([]);
const logInfo = ref<LogInfo>({})
const loading = ref(false);
const dialogInfoVisible = ref(false)
const newVersion = ref('')

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
  await GetLogInfo().then(result => {
    logInfo.value = result
  })
  dialogInfoVisible.value = true
}

const mergeEreRecord = async (typ: string) => {
  loading.value = true
  await MergeEreRecord(currentUid.value, typ).then(() => {
    ElMessage({message: '合并成功', type: 'success', plain: true, showClose: true, duration: 2000})
  }).catch(() => {
    ElMessage({message: '合并发生错误', type: 'error', plain: true, showClose: true, duration: 2000})
  })
  await getAllPoolInfo()
  loading.value = false
}

const copyUid = () => {
  ClipboardSetText(logInfo.value.uid)
  ElMessage({message: 'UID已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

const copyGachaUrl = () => {
  ClipboardSetText(logInfo.value.gachaUrl)
  ElMessage({message: '抽卡链接已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

const copyAccessToken = () => {
  ClipboardSetText(logInfo.value.accessToken)
  ElMessage({message: 'AccessToken已复制', type: 'success', plain: true, showClose: true, duration: 1000})
}

const handleCommunityTasks = () => {
  HandleCommunityTasks().then(result => {
    ElMessage({message: result.join("<br/>"), type: 'success', plain: true, showClose: true, duration: 3000, dangerouslyUseHTMLString: true})
  }).catch(err => {
    ElMessage({message: err, type: 'error', plain: true, showClose: true, duration: 3000})
  })
}

const checkUpdate = () => {
  CheckUpdate().then(result => {
    newVersion.value = result
  })
}

const applyUpdate = async () => {
  loading.value = true
  await ApplyUpdate()
  loading.value = false
}

onMounted(async () => {
  await getUidList()
  if (uidList.value.length > 0) {
    currentUid.value = uidList.value[0]
    await getAllPoolInfo()
  }
  checkUpdate()
})

</script>

<template>
  <div class="h-dvh w-full flex flex-col p-4 gap-4" v-loading="loading" element-loading-text="Loading...">
    <div class="flex">
      <div class="grow">
        <el-button type="success" class="font-bold" @click="incrementalUpdatePoolInfo">增量更新</el-button>
        <el-button type="primary" class="font-bold" disabled>全量更新</el-button>
        <el-dropdown class="ml-3">
          <el-button type="danger" class="font-bold">导入导出</el-button>
          <template #dropdown>
            <el-dropdown-menu :disabled="!currentUid">
              <el-dropdown-item @click="mergeEreRecord('json')">导入EreJson</el-dropdown-item>
              <el-dropdown-item @click="mergeEreRecord('excel')">导入EreExcel</el-dropdown-item>
              <el-dropdown-item divided disabled>导出Json</el-dropdown-item>
              <el-dropdown-item disabled>导出Excel</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <el-button type="warning" class="font-bold ml-3" v-if="newVersion" @click="applyUpdate">更新到{{newVersion}}</el-button>
      </div>
      <div class="flex items-center gap-2">
        <el-button type="primary" class="font-bold" @click="handleCommunityTasks">一键社区</el-button>
        <div>UID:</div>
        <el-select v-model="currentUid" class="w-28" @change="getAllPoolInfo">
          <el-option v-for="uid in uidList" :key="uid" :label="uid" :value="uid"/>
        </el-select>
        <el-button text :icon="Connection" circle @click="openInfoDialog"/>
      </div>
    </div>
    <div class="w-full flex flex-wrap gap-4">
      <PoolCard v-for="pool in poolList" :pool="pool"></PoolCard>
    </div>
    <el-dialog v-model="dialogInfoVisible" width="600">
      <template #title>
        <div class="text-xl font-bold">关于</div>
      </template>
      <div class="flex flex-col gap-4">
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">项目地址</div>
          <div class="grow text-blue-500">https://github.com/MatchaCabin/gf2gacha</div>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">UID</div>
          <el-input class="grow" readonly v-model="logInfo.uid"/>
          <el-button text :icon="CopyDocument" circle @click="copyUid"/>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">抽卡链接</div>
          <el-input class="grow" readonly v-model="logInfo.gachaUrl"/>
          <el-button text :icon="CopyDocument" circle @click="copyGachaUrl"/>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-24 shrink-0">AccessToken</div>
          <el-input class="grow" readonly type="password" v-model="logInfo.accessToken"/>
          <el-button text :icon="CopyDocument" circle @click="copyAccessToken"/>
        </div>
        <el-alert title="AccessToken是您的临时登录凭证，请自行把控风险，切勿随意泄露" type="warning" show-icon :closable="false"></el-alert>
      </div>
    </el-dialog>
  </div>
</template>

<style>

</style>
