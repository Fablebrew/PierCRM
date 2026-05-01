<template>
    <Header />
    <Nav page-title="Мой тариф"/>
    <section class="tarif-page">
        <div class="page-head">
            <h1>Ваш тариф: <b>Пробный</b></h1>
            <p>Тариф оплачен до: 14.06.2026</p>
            <div class="head-buttons">
                <button class="history">История платежей</button>
                <button class="reject">Отказаться от тарифа</button>
            </div>
        </div>
        <div class="page-body">
            <h1>Тарифные планы</h1>
            <div class="cards">
                <div class="card">
                    <div class="card-text">
                        <h2>Название тарифа</h2>
                        <div class="cost"><h3>$9.9<span>/month</span></h3></div>
                        <ul>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Описание</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Component Properties</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Interactive Components</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Light & Dark Theme</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>30+ Page Examples</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Lifetime Updates</p>
                            </li>                            
                        </ul>
                    </div>
                    <button>Выбрать тариф</button>
                </div>
                <div class="card">
                    <div class="card-text">
                        <h2>Название тарифа</h2>
                        <div class="cost"><h3>$9.9<span>/month</span></h3></div>
                        <ul>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Описание</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Component Properties</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Interactive Components</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Light & Dark Theme</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>30+ Page Examples</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Lifetime Updates</p>
                            </li>    
                        </ul>
                    </div>
                    <button>Выбрать тариф</button>
                </div>
                <div class="card">
                    <div class="card-text">
                        <h2>Название тарифа</h2>
                        <div class="cost"><h3>$9.9<span>/month</span></h3></div>
                        <ul>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Описание</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Component Properties</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Interactive Components</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Light & Dark Theme</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>30+ Page Examples</p>
                            </li>
                            <li>
                                <img src="@/assets/img/IconSet.svg" alt="set_icon">
                                <p>Lifetime Updates</p>
                            </li>    
                        </ul>
                    </div>
                    <button>Выбрать тариф</button>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import Header from './Header.vue'
import Nav from './Nav.vue'
export default {
    components: {
    Header,
    Nav
  },
}
</script>

<style scoped>
.tarif-page {
    background-color: #F1F4FF;
    margin: 0%;
    padding: 20px;
}
.page-head {
    text-align: center;
}
.page-head h1 {
    font-family: "Manrope";
    font-weight: 400;
    font-style: "Regular";
    font-size: 30px;
    line-height: 140%;
    letter-spacing: 0%;
}
.page-head p {
    font-family: "Manrope";
    font-weight: 400;
    font-style: "Regular";
    font-size: 22px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
}
.head-buttons button {
    border-radius: 4px;
    outline: none;
    padding: 14px 40px 14px 40px;
}
.history {
    margin-right: 7.5px;
    color: white;
    background-color: #3383FB;
    border: none;
}
.reject {
    margin-left: 7.5px;
    background-color: #FFFFFF;
    border: 0.5px solid #BBBBBB;
}
.page-body {
    text-align: center;
}
.page-body h1 {
    font-family: "Manrope";
    font-weight: 500;
    font-style: "Medium";
    font-size: 30px;
    line-height: 140%;
    letter-spacing: 0%;
}
.cards {
    display: flex;
    justify-content: center;
}
.card {
    width: 33%;
    background-color: #DBE1F675;
    transition: 0.3s;
}
.card:hover {
    background-color: #FFFFFF;
}
.card-text {
    text-align: left;
}
.card-text h2 {
    font-family: "Inter";
    font-weight: 600;
    font-style: "Semi Bold";
    font-size: 18px;
    line-height: 28px;
    letter-spacing: 0%;
}
.cost h3 {
    font-family: "Inter";
    font-weight: 600;
    font-style: "Semi Bold";
    font-size: 48px;
    line-height: 58px;
    letter-spacing: 0%;
    color: #000000B2;
}
.cost span {
    font-family: "Inter";
    font-weight: 400;
    font-style: "Regular";
    font-size: 14px;
    line-height: 20px;
    letter-spacing: 0%;
    color: #00000066;
}
.card-text ul {
    list-style: none;
}
.card-text ul li {
    display: flex;
    align-items: center;
    gap: 8px; /* отступ между иконкой и текстом */
}
.card-text ul li p {
    font-family: "Inter";
    font-weight: 400;
    font-style: "Regular";
    font-size: 18px;
    line-height: 28px;
    letter-spacing: 0%;
    color: #000000B2;
    margin: 0; /* убираем стандартные отступы */
}
.card button {
    width: 100%;
    background-color:#3383FB;
    color: #FFFFFF;
    font-family: "Inter";
    font-weight: 400;
    font-style: "Regular";
    font-size: 14px;
    line-height: 28px;
    letter-spacing: 0%;
    text-align: center;
    vertical-align: middle;
    border: none;
    border-radius: 4px;
    padding: 10px;
}
</style>