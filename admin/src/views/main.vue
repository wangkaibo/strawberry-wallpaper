<template>
<div class="sw-main">
    <h3>草莓壁纸-Strawberry Wallpaper 后台统计平台</h3>
    <div class="sw-row">
        <div class="sw-main_data">
            <div class="sw-main_title">安装数量统计</div>
            <div class="sw-main_data_title">
                <span>安装总量:</span><span>{{ totalNum }}</span>
            </div>
            <div class="sw-main_data_title">
                <span>最新的活跃量:</span><span>{{ activeNum }}</span>
            </div>
        </div>
        <div class="sw-main_system">
            <div class="sw-main_title">安装平台统计</div>
            <div class="sw-main_system_content">
                <div ref="ref_system" class="sw-main_system_platform"></div>
                <div class="sw-main_system_version">
                    <div class="sw-main_system_version_mac sw-main_system_version_item">
                        <div>mac</div>
                           <div><span>10.10.14</span>: <span>12</span></div> 
                           <div><span>10.10.14</span>: <span>12</span></div> 
                           <div><span>10.10.14</span>: <span>12</span></div> 
                    </div>
                    <div class="sw-main_system_version_win sw-main_system_version_item">
                        <div>win</div>
                        <div><span>win7</span>: <span>12</span></div> 
                        <div><span>win8</span>: <span>12</span></div> 
                        <div><span>win10</span>: <span>12</span></div> 
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="sw-row sw-main_add">
        <div class="sw-main_title">每日新增用户量统计</div>
        <div ref="ref_add" class="sw-main_add_content sw-row_content"></div>
    </div>
    <div class="sw-row sw-main_active">
        <div class="sw-main_title">每日活跃用户量统计</div>
        <div ref="ref_active" class="sw-main_active_contentm sw-row_content"></div>
    </div>
</div>
</template>

<script>
// import * as echarts from 'echarts/lib/echarts';
import echarts from 'echarts'
import { apiStatistic } from './api.js';


// 引入 echarts 主模块。
// 引入折线图。
// import 'echarts/lib/chart/line';
// 引入饼图
// import 'echarts/lib/chart/pie';
// 引入提示框组件、标题组件、工具箱组件。
// import 'echarts/lib/component/tooltip';
// import 'echarts/lib/component/title';
// import 'echarts/lib/component/toolbox';

let echartsSystem = null;
let echartsAdd = null;
let echartsActive = null;

export default {
    name: 'swMain',

    data() {
        return {
            totalNum: '12',
            activeNum: '12',
        };
    },
    created() {},
    mounted() {
        this.$nextTick(() => {
            // this.echartsSystemInit({
            //     mac: 35,
            //     win: 40
            // });
            // this.echartsAddInit([
            //     { date: '2018-05-01', num: 34 },
            //     { date: '2018-05-02', num: 56 },
            //     { date: '2018-05-03', num: 23 },
            //     { date: '2018-05-04', num: 38 },
            //     { date: '2018-05-05', num: 30 },
            //     { date: '2018-05-06', num: 67 },
            //     { date: '2018-05-07', num: 35 },
            //     { date: '2018-05-08', num: 6 },
            //     { date: '2018-05-09', num: 56 },
            // ]);

            // this.echartsActiveInit([
            //     { date: '2018-05-01', num: 34 },
            //     { date: '2018-05-02', num: 56 },
            //     { date: '2018-05-03', num: 23 },
            //     { date: '2018-05-04', num: 38 },
            //     { date: '2018-05-05', num: 30 },
            //     { date: '2018-05-06', num: 67 },
            //     { date: '2018-05-07', num: 35 },
            //     { date: '2018-05-08', num: 6 },
            //     { date: '2018-05-09', num: 56 },
            // ]);
        })
      
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
            echartsSystem = echarts.init(this.$refs.ref_system);
            echartsSystem.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{a} <br/>{b}: {c} ({d}%)'
                },
                series: [
                    {
                        name: '安装平台',
                        type: 'pie',
                        radius: ['50%', '70%'],
                        avoidLabelOverlap: false,
                        label: {
                            normal: {
                                show: true,
                                fontSize: 20,
                                formatter: '{b}:{c}台\n{d}%',
                                textStyle: {
                                    fontSize: '20',
                                    fontWeight: 'bold'
                                }
                            },
                            emphasis: {
                                show: true,
                                position: 'center',
                                textStyle: {
                                    fontSize: '30',
                                    fontWeight: 'bold'
                                }
                            }
                        },
                        labelLine: {
                            show: true,
                        },
                        data: resultData
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
            echartsAdd = echarts.init(this.$refs.ref_add);
            echartsAdd.setOption({
                grid: {
                    left: 60,
                    right: 60,
                    bottom: 80,
                    top: 20,
                },
                tooltip: {
                    trigger: 'axis',
                    position(pt) {
                        return [pt[0], '10%'];
                    }
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: xData,
                    axisLine: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisTick: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisLabel: {
                        color: '#fff'
                    },
                    splitLine: {
                        lineStyle: {
                            // 使用深浅的间隔色
                            color: ['#aaa', '#ddd']
                        }
                    }
                },
                yAxis: {
                    type: 'value',
                    axisLine: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisTick: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisLabel: {
                        color: '#fff'
                    }, 
                    splitLine: {
                        lineStyle: {
                            // 使用深浅的间隔色
                            color: ['#aaa', '#ddd']
                        }
                    },
                    splitArea: {
                        interval: 1
                    }

                },
                dataZoom: [{
                    type: 'slider',
                    start: 0,
                    end: 100,
                }],
                series: [
                    {
                        smooth: true,
                        symbol: 'none',
                        sampling: 'average',
                        itemStyle: {
                            color: 'rgb(255, 70, 131)'
                        },
                        data: yData,
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
            echartsActive = echarts.init(this.$refs.ref_active);
            echartsActive.setOption({
                grid: {
                    left: 60,
                    right: 60,
                    bottom: 80,
                    top: 20,
                },
                tooltip: {
                    trigger: 'axis',
                    position(pt) {
                        return [pt[0], '10%'];
                    }
                },
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: xData,
                    axisLine: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisTick: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisLabel: {
                        color: '#fff'
                    }
                },
                yAxis: {
                    type: 'value',
                    axisLine: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisTick: {
                        lineStyle: {
                            color: '#fff'
                        }
                    },
                    axisLabel: {
                        color: '#fff'
                    }

                },
                dataZoom: [{
                    type: 'slider',
                    start: 0,
                    end: 100,
                }],
                series: [
                    {
                        smooth: true,
                        symbol: 'none',
                        sampling: 'average',
                        itemStyle: {
                            color: 'rgb(255, 70, 131)'
                        },
                        data: yData,
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
