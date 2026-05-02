<template>
    <Header />
    <Nav page-title="Мой тариф" class="nav-component"/>
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
            <div class="cards-container">
                <div class="cards">
                    <div class="card tarif-card">
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
                    <div class="card tarif-card">
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
                    <div class="card tarif-card">
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
* {
    box-sizing: border-box;
}
.tarif-card {
    margin: 0 !important;
    margin-left: 0 !important;
    margin-right: 0 !important;
    /* остальные стили */
}
.tarif-page {
    background-color: #F1F4FF;
    margin: 0;
    padding: 20px 0;
    overflow-x: hidden;
    width: 100%;
    min-height: 100vh;
}

.page-head {
    text-align: center;
    max-width: 100%;
    margin: 0 auto;
    padding: 0 20px;
}

.page-head h1 {
    font-family: "Manrope";
    font-weight: 400;
    font-style: "Regular";
    font-size: 30px;
    line-height: 140%;
    letter-spacing: 0%;
    word-wrap: break-word;
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

.head-buttons {
    display: flex;
    justify-content: center;
    gap: 15px;
    flex-wrap: wrap;
}

.head-buttons button {
    border-radius: 4px;
    outline: none;
    padding: 14px 40px;
    cursor: pointer;
    font-family: "Manrope";
    font-weight: 400;
    font-size: 16px;
    white-space: nowrap;
}

.history {
    color: white;
    background-color: #3383FB;
    border: none;
}

.reject {
    background-color: #FFFFFF;
    border: 0.5px solid #BBBBBB;
}

.page-body {
    width: 100%;
    margin: 0 auto;
    padding: 0 20px;
}

.page-body h1 {
    font-family: "Manrope";
    font-weight: 500;
    font-style: "Medium";
    font-size: 30px;
    line-height: 140%;
    letter-spacing: 0%;
    text-align: center;
}

.cards-container {
    width: 100%;
    display: flex;
    justify-content: center;
}

.cards {
    display: flex;
    justify-content: center;
    gap: 20px;
    width: 100%;
    max-width: 1200px;
}

.card {
    flex: 1;
    width: 33.333%;
    background-color: #DBE1F675;
    transition: 0.3s;
    border-radius: 8px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.card:hover {
    background-color: #FFFFFF;
}

.card-text {
    text-align: left;
    flex: 1;
}

.card-text h2 {
    font-family: "Inter";
    font-weight: 600;
    font-style: "Semi Bold";
    font-size: 18px;
    line-height: 28px;
    letter-spacing: 0%;
    word-wrap: break-word;
}

.cost h3 {
    font-family: "Inter";
    font-weight: 600;
    font-style: "Semi Bold";
    font-size: 48px;
    line-height: 58px;
    letter-spacing: 0%;
    color: #000000B2;
    word-wrap: break-word;
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
    padding: 0;
}

.card-text ul li {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;
}

.card-text ul li p {
    font-family: "Inter";
    font-weight: 400;
    font-style: "Regular";
    font-size: 18px;
    line-height: 28px;
    letter-spacing: 0%;
    color: #000000B2;
    margin: 0;
    flex: 1;
}

.card button {
    width: 100%;
    background-color: #3383FB;
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
    cursor: pointer;
    margin-top: 20px;
}

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
    .nav-component {
        display: none !important;
    }
    
    .page-head {
        padding: 0 15px;
    }
    
    .page-head h1 {
        font-size: 22px;
    }
    
    .page-head p {
        font-size: 16px;
    }
    
    .head-buttons {
        flex-direction: column;
        align-items: center;
    }
    
    .head-buttons button {
        padding: 12px 24px;
        font-size: 14px;
        width: 100%;
        max-width: 300px;
    }
    
    .page-body {
        padding: 0 15px;
    }
    
    .page-body h1 {
        font-size: 22px;
    }
    
    .cards {
        flex-direction: column;
        align-items: center;
        gap: 20px;
    }
    
    .card {
        width: 100%;
        max-width: 400px;
        flex: none;
    }
    
    .card-text h2 {
        font-size: 16px;
    }
    
    .cost h3 {
        font-size: 36px;
        line-height: 44px;
    }
    
    .card-text ul li p {
        font-size: 16px;
        line-height: 24px;
    }
}

@media (max-width: 480px) {
    .tarif-page {
        padding: 15px 0;
    }
    
    .page-head {
        padding: 0 10px;
    }
    
    .page-head h1 {
        font-size: 18px;
    }
    
    .page-head p {
        font-size: 14px;
    }
    
    .head-buttons button {
        padding: 10px 20px;
        font-size: 12px;
    }
    
    .page-body {
        padding: 0 10px;
    }
    
    .card {
        padding: 15px;
        max-width: 100%;
    }
    
    .cost h3 {
        font-size: 30px;
        line-height: 38px;
    }
    
    .card-text ul li p {
        font-size: 14px;
        line-height: 20px;
    }
    
    .card button {
        font-size: 12px;
        padding: 8px;
    }
}
</style>