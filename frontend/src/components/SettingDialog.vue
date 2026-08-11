<script setup lang="ts">
import {onBeforeMount, ref} from "vue";
import {GetCommunityExchangeList, GetSettingCapturePort, GetSettingExchangeList, GetSettingFont, GetSettingLayout, SaveSettingCapturePort, SaveSettingExchangeList, SaveSettingFont, SaveSettingLayout} from "../../wailsjs/go/main/App";
import {model} from "../../wailsjs/go/models.ts";
import {ElMessage} from "element-plus";
import {useLayoutStore} from "../stores/layout.ts";
import CommunityExchangeList = model.CommunityExchangeList;

const visible = defineModel({type: Boolean, required: true})

const exchangeList = ref<CommunityExchangeList[]>([])
const exchangeSelectedList = ref<number[]>([])
const onExchangeListChange = () => {
  SaveSettingExchangeList(exchangeSelectedList.value)
}

const selectedFont = ref('')
const fontList = ref<{ name: string, value: string }[]>([])
const onFontChange = async (newFont: string) => {
  await SaveSettingFont(newFont)
  let body = document?.querySelector('body')
  if (body) {
    body.style.fontFamily = newFont
  }
}

const layoutStore = useLayoutStore()
const onLayoutChange = async (newLayoutType: number) => {
  await SaveSettingLayout(newLayoutType)
}

const capturePort = ref<number | null>(null)
const onPortChange = async (newPort: number) => {
  await SaveSettingCapturePort(newPort)
}

onBeforeMount(async () => {
  await GetCommunityExchangeList().then(result => {
    exchangeList.value = result
  })
  await GetSettingExchangeList().then(result => {
    if (result) {
      exchangeSelectedList.value = result
    }
  }).catch(err => {
    ElMessage({message: err, type: 'error', plain: true, showClose: true, duration: 0})
  })
  //@ts-ignore
  await queryLocalFonts().then(result => {
    if (result) {
      //@ts-ignore
      result.forEach(item => {
        fontList.value.push({name: item.fullName, value: item.fullName})
      })
    }
  })
  await GetSettingFont().then(result => {
    if (result) {
      selectedFont.value = result
    } else {
      selectedFont.value = 'ChillRoundM'
    }
  })

  await GetSettingLayout().then(result => {
    if (result) {
      layoutStore.layoutType = result
    }
  })

  await GetSettingCapturePort().then(result => {
    if (result) {
      capturePort.value = result
    }
  })
})

</script>

<template>
  <el-dialog v-model="visible" width="600" destroy-on-close>
    <template #header>
      <div class="text-xl font-bold">设置</div>
    </template>
    <div class="flex flex-col gap-4">
      <div class="flex items-center gap-2">
        <div class="w-24 shrink-0">社区兑换</div>
        <div class="grow text-blue-500">
          <el-checkbox-group v-model="exchangeSelectedList" @change="onExchangeListChange" class="flex flex-wrap">
            <el-checkbox class="basis-1/2 !m-0" v-for="item in exchangeList" :label="item.name" :value="item.id"></el-checkbox>
          </el-checkbox-group>
        </div>
      </div>
      <div class="flex items-center gap-2">
        <div class="w-24 shrink-0">字体</div>
        <el-select class="font-serif" v-model="selectedFont" @change="onFontChange" filterable>
          <el-option class="font-serif" key="ChillRoundM" label="寒蝉半圆体" value="ChillRoundM"/>
          <el-option class="font-serif" v-for="font in fontList" :key="font.value" :label="font.name" :value="font.value"/>
        </el-select>
      </div>
      <div class="flex items-center gap-2">
        <div class="w-24 shrink-0">布局</div>
        <el-select v-model="layoutStore.layoutType" @change="onLayoutChange">
          <el-option :key="0" label="宽松" :value="0"/>
          <el-option :key="1" label="紧凑" :value="1"/>
        </el-select>
      </div>
      <div class="flex items-center gap-2">
        <div class="w-24 shrink-0">捕获代理端口</div>
        <el-input-number
          v-model="capturePort"
          :min="1"
          :max="65535"
          placeholder="留空则自动分配"
          @change="onPortChange"
          controls-position="right"
          style="width: 50%"
        />
      </div>
    </div>
  </el-dialog>
</template>