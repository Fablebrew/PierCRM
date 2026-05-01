<template>
    <header>
    <div class="header-content">
        <img class="search-icon" src="@/assets/img/Union.svg" alt="search_icon">
        <div class="header-left">
            <h1 class="logo">Vedi Project</h1>
            <div class="left-item">
                <input type="text" class="search-input" placeholder="Поиск">
            </div>
        </div>
        <div class="header-right">
            <a href="/tarif" class="header-nav">Мой тариф</a>
            <a href="#" class="header-nav">Партнерская программа</a>
            <button>I</button>
            <img class="avatar" src="@/assets/img/user_photo.png" alt="avatar">
        </div>
    </div>
</header>
</template>
<script>
export default {
    name: "Header",
}
</script>

<style scoped>
header {
    padding: 2% 5% 2% 5%;
    border-bottom: 1px solid #BBBBBB;
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
.logo {
    font-family: "Manrope";
    font-weight: 600;
    font-size: 20px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #3383FB;
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
.left-item {
    align-content: center;
}
</style>