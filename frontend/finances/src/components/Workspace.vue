<template>
<Header></Header>
<Nav pageTitle="Рабочий стол"></Nav>
<section id="main-content">
    <div class="cards">
        <div class="row">
            <div class="row-left">
                <div class="balance card">
                    <h1>Денежный баланс</h1>
                    <span>86 500 000р.</span>
                    <p>Весь период</p>
                </div>
                <div class="zp card">
                    <div class="card-head">
                        <div class="card-head_title">
                            <h1>Выручка</h1>
                        </div>
                        <div class="card-head_data">
                            <p>133 506 000р.</p>
                        </div>
                        <div class="card-head_buttons">
                            <button class="head-button">Ы</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="dynamics card">
                <div class="card-head">
                    <div class="card-head_title">
                        <h1>Динамика выручки и расходов по месяцам</h1>
                    </div>
                    <div class="card-head_buttons">
                        <button class="head-button">Год</button>
                        <button class="head-button">12 месяцев</button>
                        <button class="head-button">Квартал</button>
                        <button class="head-button">Ы</button>
                    </div>
                </div>
                <!-- <div class="dynamics-chart"></div> -->
                <DynamicChart :chartData="data" />
            </div>
        </div>
        <div class="row">
            <div class="other card">
                <div class="row">
                    <div class="card-text">
                        <h2>Расходы по проектам</h2>
                        <p>33 506 000р.</p>
                    </div>
                    <div class="circle-chart">
                        <DonutChart />
                    </div>
                    <div class="card-butt">
                        <button></button>
                    </div>
                </div>
            </div>
            <div class="other card">
                <div class="row">
                    <div class="card-text">
                        <h2>Расходы по проектам</h2>
                        <p>33 506 000р.</p>
                    </div>
                    <div class="circle-chart">
                        <DonutChart />
                    </div>
                    <div class="card-butt">
                        <button></button>
                    </div>
                </div>
            </div>
            <div class="other card">
                <div class="row">
                    <div class="card-text">
                        <h2>Расходы по проектам</h2>
                        <p>33 506 000р.</p>
                    </div>
                    <div class="circle-chart">
                        <DonutChart />
                    </div>
                    <div class="card-butt">
                        <button></button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>
<div class="bottom-button">
        <button><a href="/projects">Перейти к проектам</a></button>
    </div>
</template>

<script>
import DynamicChart from "@/components/DynamicChart.vue";
import Revenue from "@/components/Revenue.vue";
import DonutChart from "@/components/DonutChart.vue";
import Header from "./Header.vue";
import Nav from "./Nav.vue";


export default {
  components: 
    { 
        DynamicChart,
        Revenue,
        DonutChart,
        Header,
        Nav
    },
  data() {
    return {
      data: [
        { label: "Янв", expenses: 70, income: 90 },
        { label: "Фев", expenses: 50, income: 80 },
        { label: "Мар", expenses: 90, income: 100 },
        { label: "Апр", expenses: 30, income: 60 },
        { label: "Май", expenses: 80, income: 120 },
      ],
    //   revdata: [
    //     { label: "Янв", value: 40 },
    //     { label: "Фев", value: 60 },
    //     { label: "Мар", value: 80 },
    //     { label: "Апр", value: 50 },
    //     { label: "Май", value: 90 },
    //   ],
    };
  },
};


</script>

<style>

.other {
    background-color: #FFFFFF;
    border-radius: 20px;
    margin-top: 1%;
}
.bottom-button {
    width: 100%;
}
.bottom-button button {
    font-family: "Manrope";
    font-weight: 400;
    font-style: "Regular";
    font-size: 22px;
    line-height: 150%;
    letter-spacing: 0%;
    background-color:#446BFA;
    color: #FFFFFF;
    outline: none;
    border: none;
    border-radius: 16px;
    width: 100%;
    padding: 20px;
}
.bottom-button button a {
    text-decoration: none;
    color: #FFFFFF;
}
.bottom-button button:hover {
    cursor: pointer;
}
.zp {
    background-color: #FFFFFF;
}
.row {
    display: flex;
    /* justify-content: space-around; */
}
.card {
    padding: 20px;
    box-shadow: 5px 5px 10px rgba(197, 197, 197, 0.692);
}
.card-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
}
.card-head_title h1 {
    font-family: "Manrope";
    font-weight: 400;
    font-style: Regular;
    font-size: 22px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #000000B2;
}
.card-head_data p {
    font-family: "Manrope";
    font-weight: 500;
    font-style: Medium;
    font-size: 44px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #000000B2;
}
#main-content {
    /* background-color: rgb(243, 243, 243); */
    background-color: #F1F4FF;
    padding: 2% 5% 2% 5%;
}
.dynamics {
    background-color: #FFFFFF;
    width: 65%;
}
.balance {
    text-align: center;
    background: linear-gradient(90deg, #00C6FB 0%, #446BFA 100%);
    padding: 2% 5% 2% 5%;
}
.balance h1 {
    font-family: "Manrope";
    font-weight: 400;
    font-style: Regular;
    font-size: 22px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #FFFFFF;
}
.balance span {
    font-family: "Manrope";
    font-weight: 500;
    font-style: Medium;
    font-size: 44px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #FFFFFF;
}
.balance p {
    font-family: "Manrope";
    font-weight: 400;
    font-style: Regular;
    font-size: 18px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #FFFFFF;
}
.cards {
    /* display: flex; */
}
.card {
    border-radius: 20px;
    margin-left: 20px;
}

.row-left {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 35%;
}
.row-left .card {
    flex: 1;
}


</style>