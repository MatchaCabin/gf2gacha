<script setup lang="ts">
import {model} from "../../wailsjs/go/models";
import {use} from 'echarts/core';
import {PieChart} from 'echarts/charts';
import {LegendComponent, TitleComponent, TooltipComponent} from 'echarts/components';
import {CanvasRenderer} from 'echarts/renderers';
import VChart from 'vue-echarts';
import Pool = model.Pool;

use([PieChart, TitleComponent, TooltipComponent, LegendComponent, CanvasRenderer]);

const props = defineProps<{ pool: Pool }>()
const title = (): string => {
  switch (props.pool.PoolType) {
    case 1:
      return '常驻池'
    case 3:
      return '角色池'
    case 4:
      return '武器池'
    case 5:
      return '新手池'
    case 8:
      return '神秘箱'
    default:
      return `未知PoolType: ${props.pool.PoolType}`
  }
}
const tagType = ['primary', 'success', 'warning', 'danger']

const option = {
  tooltip: {
    trigger: 'item',
  },
  legend: {
    orient: 'horizontal',
    left: 'center',
  },
  series: [
    {
      type: 'pie',
      radius: '80%',
      data: [
        {value: props.pool.Rank5Count, name: '五星', itemStyle: {color: '#fdcb51'}},
        {value: props.pool.Rank4Count, name: '四星', itemStyle: {color: '#ddb0e2'}},
        {value: props.pool.Rank3Count, name: '三星', itemStyle: {color: '#409EFF'}}],
      itemStyle: {
        borderRadius: 4, // 让每个扇形区域有圆角
        borderColor: '#fff',
        borderWidth: 2,
      },
      label: {
        show: false,
      },
      emphasis: {
        itemStyle: {
          shadowBlur: 20,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)',
        },
      },
    },
  ],
}

</script>

<template>
  <div class="basis-72 shrink-0 grow-0 flex flex-col items-center gap-2 my-2" v-if="pool.GachaCount">
    <div class="font-bold text-xl">{{ title() }}</div>
    <v-chart class="h-64" :option="option"></v-chart>
    <div class="flex flex-col w-full">
      <div class="w-full text-sm">一共 <span class="text-red-600">{{ pool.GachaCount }}</span> 抽，已垫 <span class="text-red-600">{{ pool.StoredCount }}</span> 抽</div>
      <div class="w-full text-sm text-amber-600">五星: {{ pool.Rank5Count }} [{{ pool.GachaCount > 0 ? Math.round(pool.Rank5Count * 10000 / pool.GachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-purple-600">四星: {{ pool.Rank4Count }} [{{ pool.GachaCount > 0 ? Math.round(pool.Rank4Count * 10000 / pool.GachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-blue-600">三星: {{ pool.Rank3Count }} [{{ pool.GachaCount > 0 ? Math.round(pool.Rank3Count * 10000 / pool.GachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-green-600">平均出金抽数：{{ pool.Rank5Count > 0 ? (pool.GachaCount / pool.Rank5Count).toFixed(1)  : '无' }}</div>
      <div class="w-full text-sm text-red-600" v-if="pool.PoolType==3||pool.PoolType==4">歪率: {{ pool.Rank5Count > 0 ?Math.round(pool.LoseCount * 10000 / (pool.Rank5Count - pool.GuaranteesCount)) / 100 + '%' : '0%'}} </div>
    </div>
    <div class="w-full text-sm text-gray-400">五星抽卡记录：</div>
    <div class="flex flex-wrap gap-1">
      <el-tag class="text-sm" v-for="record in pool.RecordList" effect="dark" :type="tagType[record.Id%4]">{{ `${record.Name}(${record.Count}${record.Lose ? '歪' : ''})` }}</el-tag>
    </div>
  </div>
</template>