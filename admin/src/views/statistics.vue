<template>
    <div class="sw-main-body">
        <div class="sw-main">
            <div class="sw-main_header">
                <div class="sw-main_header_h3">Strawberry Wallpaper 后台统计平台</div>
                <div class="sw-main_header_icon"></div>
                <div class="sw-main_header_date">{{currentDate}}</div>
            </div>
            <div class="sw-row sw-row_one">
                <div class="sw-main_data">
                    <div class="sw-modal-bk"></div>
                    <div class="sw-modal-content">
                        <div class="sw-main_title">安装数量统计</div>
                        <div class="sw-main_data_title">
                            <div class="sw-main_data_title_t">安装总量:</div>
                            <div class="sw-main_data_title_c">{{ totalNum }}</div>
                        </div>
                        <div class="sw-main_data_title">
                            <div class="sw-main_data_title_t">最新的日活量:</div>
                            <div class="sw-main_data_title_c">{{ activeNum }}</div>
                        </div>
                    </div>
                </div>
                <div class="sw-main_system">
                    <div class="sw-modal-bk"></div>
                    <div class="sw-modal-content">
                        <div class="sw-main_title">安装平台统计</div>
                        <div class="sw-main_system_content">
                            <div ref="ref_system" class="sw-main_system_platform"></div>
                            <!-- <div class="sw-main_system_version">
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
                        </div> -->
                        </div>
                    </div>
                </div>
            </div>
            <div class="sw-row sw-main_add">
                <div class="sw-modal-bk"></div>
                <div class="sw-modal-content">
                    <div class="sw-main_title">每日新增用户量统计</div>
                    <div ref="ref_add" class="sw-main_add_content sw-row_content"></div>
                </div>
            </div>
            <div class="sw-row sw-main_active">
                <div class="sw-modal-bk"></div>
                <div class="sw-modal-content">
                    <div class="sw-main_title">每日活跃用户量统计</div>
                    <div ref="ref_active" class="sw-main_active_contentm sw-row_content"></div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
// import * as echarts from 'echarts/lib/echarts';
import echarts from 'echarts'
import { apiStatistic } from './api.js'
import { dateFormat } from '../assets/js/until.js'

// 引入 echarts 主模块。
// 引入折线图。
// import 'echarts/lib/chart/line';
// 引入饼图
// import 'echarts/lib/chart/pie';
// 引入提示框组件、标题组件、工具箱组件。
// import 'echarts/lib/component/tooltip';
// import 'echarts/lib/component/title';
// import 'echarts/lib/component/toolbox';

let echartsSystem = null
let echartsAdd = null
let echartsActive = null
const lineConfig = (xData, yData) => ({
    grid: {
        left: 60,
        right: 60,
        bottom: 80,
        top: 20,
    },
    tooltip: {
        trigger: 'axis',
        position(pt) {
            return [pt[0], '10%']
        }
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: xData,
        axisLine: {
            lineStyle: {
                color: '#aaa'
            }
        },
        axisTick: {
            lineStyle: {
                color: '#aaa'
            }
        },
        axisLabel: {
            color: '#aaa'
        }
    },
    yAxis: {
        type: 'value',
        axisLine: {
            lineStyle: {
                color: '#aaa'
            }
        },
        axisTick: {
            lineStyle: {
                color: '#aaa'
            }
        },
        axisLabel: {
            color: '#aaa'
        }

    },
    dataZoom: [{
        type: 'slider',
        start: 0,
        end: 100,
    }],
    series: [
        {
            // smooth: true,
            symbol: 'none',
            sampling: 'average',
            itemStyle: {
                color: '#2cade2'
            },
            data: yData,
            type: 'line'
        }
    ]
})

const pieConfig = resultData => ({
    color: ['#24feb4', '#005eff'],
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
})

const platform = {
    Macintosh: 'Mac',
    Windows: 'Win'
}

export default {
    name: 'swMain',

    data() {
        return {
            totalNum: '0',
            activeNum: '0',
            currentDate: '',
        }
    },
    created() {
        window.setInterval(() => {
            this.currentDate = dateFormat(new Date(), 'YYYY-MM-dd hh:mm:ss')
        }, 1000)
    },
    mounted() {
        apiStatistic().then((result) => {
            this.echartsSystemInit(result.platform)
            this.echartsAddInit(result.register)
            this.echartsActiveInit(result.active)
            this.totalNum = result.total_num
            this.activeNum = result.active_num
        }).catch((e) => {
            console.log(e)
        })
    },
    methods: {
        echartsSystemInit(data) {
            const resultData = data.map(item => ({
                value: item.count,
                name: platform[item.platform] || item.platform
            }))
            if (echartsSystem === null){
                echartsSystem = echarts.init(this.$refs.ref_system)
            }
            echartsSystem.setOption(pieConfig(resultData))
        },
        echartsAddInit(data) {
            const xData = []
            const yData = []
            data.forEach((item) => {
                xData.push(item.date)
                yData.push(item.count)
            })
            if (echartsAdd === null){
                echartsAdd = echarts.init(this.$refs.ref_add)
            }
            echartsAdd.setOption(lineConfig(xData, yData))
        },
        echartsActiveInit(data) {
            const xData = []
            const yData = []
            data.forEach((item) => {
                xData.push(item.date)
                yData.push(item.count)
            })
            if (echartsActive === null){
                echartsActive = echarts.init(this.$refs.ref_active)
            }
            echartsActive.setOption(lineConfig(xData, yData))
        }
    }
}
</script>

<style lang="scss" scoped>
</style>
