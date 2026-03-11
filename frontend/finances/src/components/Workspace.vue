<template>
<header>
    <div class="header-content">
        <img class="search-icon" src="@/assets/img/union.svg" alt="search_icon">
        <div class="header-left">
            <h1 class="logo">Vedi Project</h1>
            <div class="left-item">
                <input type="text" class="search-input" placeholder="Поиск">
            </div>
        </div>
        <div class="header-right">
            <a href="#" class="header-nav">Мой тариф</a>
            <a href="#" class="header-nav">Партнерская программа</a>
            <button>I</button>
            <img class="avatar" src="@/assets/img/user_photo.png" alt="avatar">
        </div>
    </div>
</header>
<section id="nav-section">
    <div class="breadcrumb">
        <h1>Рабочий стол</h1>
    </div>
    <div class="nav">
        <ul>
            <li><a href="#">Проекты</a></li>
            <li><a href="#">Прочие расходы</a></li>
            <li><a href="#">Команда</a></li>
        </ul> 
    </div>
</section>
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
    </div>
</section>
</template>

<script>
import DynamicChart from "@/components/DynamicChart.vue";
import Revenue from "@/components/Revenue.vue";


export default {
  components: 
    { 
        DynamicChart,
        Revenue,
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
header {
    padding: 2% 5% 2% 5%;
    border-bottom: 1px solid #BBBBBB;
}
.zp {
    background-color: #FFFFFF;
}
.row {
    display: flex;
    justify-content: space-around;
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
    background-color: rgb(243, 243, 243);
    padding: 2% 5% 2% 5%;
}
.dynamics {
    background-color: #FFFFFF;
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
#nav-section {
    border-bottom: 1px solid #BBBBBB;
    display: flex;
    padding: 0.5% 5% 0.5% 5%;
    align-items: center;
    position: relative;
    background-color: rgb(243, 243, 243);
}
.breadcrumb h1 {
    font-family: "Manrope";
    font-weight: 400;
    font-style: Regular;
    font-size: 22px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #000000B2;
}
.nav {
    align-content: center;
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}
.nav ul {
    list-style: none;
    display: flex;
}
.nav ul li {
    margin-left: 20px;
}
.nav ul li a {
    text-decoration: none;
    font-family: "Manrope";
    font-weight: 400;
    font-style: Regular;
    font-size: 18px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #000000B2;
}
.logo {
    font-family: "Manrope";
    font-weight: 600;
    font-size: 20px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #3383FB;
}
.search-input {
    border-radius: 20px;
    margin-left: 20px;
    height: 45%;
    border: 1px solid #00000080;
    padding: 2% 8% 2% 16%;
    background-color: #F1F4FF;
    color: #BBBBBB;
}
.search-icon {
    position: absolute;
    left: 19.5%;
}
.header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.avatar {
    width: 64px;
    border-radius: 50%;
}
.header-left {
    display: flex;
    justify-content: space-between;
    /* width: 50%; */
}

.header-right {
    display: flex;
    justify-content: space-between;
    width: 50%;
    align-items: center;
}
.left-item {
    align-content: center;
}
.header-nav {
    text-decoration: none;
    font-family: "Manrope";
    font-weight: 400;
    font-size: 16px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
    color: #00000080;
}
</style>