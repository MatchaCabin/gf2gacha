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
  switch (props.pool.poolType) {
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
      return `未知PoolType: ${props.pool.poolType}`
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
        {value: props.pool.rank5Count, name: '五星', itemStyle: {color: '#fdcb51'}},
        {value: props.pool.rank4Count, name: '四星', itemStyle: {color: '#ddb0e2'}},
        {value: props.pool.rank3Count, name: '三星', itemStyle: {color: '#409EFF'}}],
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
  <div class="basis-72 shrink-0 grow-0 flex flex-col items-center gap-2 my-2" v-if="pool.gachaCount">
    <div class="font-bold text-xl">{{ title() }}</div>
    <div class="h-64 w-64">
      <v-chart class="h-64" :option="option"></v-chart>
    </div>
    <div class="flex flex-col w-full">
      <div class="w-full text-sm">一共 <span class="text-red-600">{{ pool.gachaCount }}</span> 抽，已垫 <span class="text-red-600">{{ pool.storedCount }}</span> 抽</div>
      <div class="w-full text-sm text-amber-600">五星: {{ pool.rank5Count }} [{{ pool.gachaCount > 0 ? Math.round(pool.rank5Count * 10000 / pool.gachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-purple-600">四星: {{ pool.rank4Count }} [{{ pool.gachaCount > 0 ? Math.round(pool.rank4Count * 10000 / pool.gachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-blue-600">三星: {{ pool.rank3Count }} [{{ pool.gachaCount > 0 ? Math.round(pool.rank3Count * 10000 / pool.gachaCount) / 100 + '%' : '0%' }}]</div>
      <div class="w-full text-sm text-green-600">平均出金抽数：{{ pool.rank5Count > 0 ? ((pool.gachaCount - pool.storedCount) / pool.rank5Count).toFixed(1) : '无' }}</div>
      <div class="w-full text-sm text-pink-400">平均出Up抽数：{{ pool.rank5Count-pool.loseCount > 0 ? ((pool.gachaCount - pool.storedCount) / (pool.rank5Count-pool.loseCount)).toFixed(1) : '无' }}</div>
      <div class="w-full text-sm text-red-600" v-if="pool.poolType==3||pool.poolType==4">歪率: {{ pool.rank5Count > 0 ? Math.round(pool.loseCount * 10000 / (pool.rank5Count - pool.guaranteesCount)) / 100 + '%' : '0%' }}</div>
    </div>
    <div class="w-full text-sm text-gray-400">五星抽卡记录：</div>
    <div class="w-full flex flex-wrap gap-2">
      <el-tag class="text-sm w-32" v-for="record in pool.recordList" effect="dark" :type="record.count<=40?'success':record.count<=60?'primary':record.count<=65?'warning':'danger'">
        <span>{{record.name}}</span>
        <span class="font-bold">「{{ record.count }}」</span>
        <span v-if="record.lose" class="text-purple-600 font-bold">歪</span>
      </el-tag>
    </div>
  </div>
</template>