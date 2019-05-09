<template>
	<div class="sw-main">
		<div class="sw-row">
			<div class="sw-main_data">
				<div>
					<span>安装总量:</span><span>{{ totalNum }}</span>
				</div>
				<div>
					<span>最新的活跃量:</span><span>{{ activeNum }}</span>
				</div>
			</div>
			<div class="sw-main_system">
				<div ref="ref_system"></div>
			</div>
		</div>
		<div class="sw-row">
			<div ref="ref_add"></div>
		</div>
		<div class="sw-row">
			<div ref="ref_active"></div>
		</div>
	</div>
</template>

<script>
import * as echarts from 'echarts/lib/echarts';
import { apiStatistic } from './api.js';

// 引入 echarts 主模块。
// 引入折线图。
import 'echarts/lib/chart/line';
// 引入饼图
import 'echarts/lib/chart/pie';
// 引入提示框组件、标题组件、工具箱组件。
import 'echarts/lib/component/tooltip';
import 'echarts/lib/component/title';
import 'echarts/lib/component/toolbox';

export default {
    name: 'main',

    data() {
        return {
            totalNum: '12',
            activeNum: '12',
            _echartsSystem: null,
            _echartsAdd: null,
            _echartsActive: null
        };
    },
    created() {},
    mounted() {
        apiStatistic().then((result) => {
            this.echartsSystemInit(result.osSta);
            this.echartsAddInit(result.installAdd);
            this.echartsActiveInit(result.activeDate);
        });
    },
    methods: {
        echartsSystemInit(data) {
            const resultData = [
                { value: data.mac, name: 'mac' },
                { value: data.win, name: 'win' }
            ];
            this._echartsSystem = echarts.init(this.$refs.ref_system);
            this._echartsSystem.setOption({
                title: {
                    text: '安装平台分类统计',
                    x: 'center'
                },
                tooltip: {
                    trigger: 'item',
                    formatter: '{a} <br/>{b} : {c} ({d}%)'
                },
                legend: {
                    type: 'scroll',
                    orient: 'vertical',
                    right: 10,
                    top: 20,
                    bottom: 20,
                    data: data.legendData,
                    selected: data.selected
                },
                series: [
                    {
                        name: '安装平台',
                        type: 'pie',
                        radius: '55%',
                        center: ['40%', '50%'],
                        data: resultData,
                        itemStyle: {
                            emphasis: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            }
                        }
                    }
                ]
            });
        },
        echartsAddInit(data) {
            const xData = [];
            const yData = [];
            data.forEach((item) => {
                xData.push(item.date);
                yData.push(item.num);
            });
            this._echartsAdd = echarts.init(this.$refs.ref_add);
            this._echartsAdd.setOption({
                xAxis: {
                    type: 'category',
                    data: yData
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: xData,
                        type: 'line'
                    }
                ]
            });
        },
        echartsActiveInit(data) {
            const xData = [];
            const yData = [];
            data.forEach((item) => {
                xData.push(item.date);
                yData.push(item.num);
            });
            this._echartsActive = echarts.init(this.$refs.ref_active);
            this._echartsActive.setOption({
                xAxis: {
                    type: 'category',
                    data: yData
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: xData,
                        type: 'line'
                    }
                ]
            });
        }
    }
};
</script>

<style lang="scss" scoped>
</style>
