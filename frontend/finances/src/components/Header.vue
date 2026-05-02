<template>
    <header>
    <div class="header-content">
        <div class="header-left">
            <h1 class="logo">Vedi Project</h1>
            <div class="search-wrapper">
                <img class="search-icon" src="@/assets/img/Union.svg" alt="search_icon">
                <input type="text" class="search-input" placeholder="Поиск">
            </div>
        </div>
        <div class="header-right">
            <a href="/tarif" class="header-nav">Мой тариф</a>
            <a href="#" class="header-nav">Партнерская программа</a>
            <button class="notification-btn">I</button>
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
* {
    box-sizing: border-box;
}

header {
    padding: 1% 3%;
    border-bottom: 1px solid #BBBBBB;
}

.header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    max-width: 1400px;
    margin: 0 auto;
    gap: 20px;
}

.header-left {
    display: flex;
    align-items: center;
    gap: 20px;
    flex: 1;
    min-width: 0;
}

.logo {
    font-family: "Manrope";
    font-weight: 600;
    font-size: clamp(16px, 2vw, 20px);
    line-height: 100%;
    color: #3383FB;
    white-space: nowrap;
    margin: 0;
}

.search-wrapper {
    position: relative;
    flex: 1;
    max-width: 400px;
    min-width: 150px;
}

.search-icon {
    position: absolute;
    left: 12px;
    top: 50%;
    transform: translateY(-50%);
    width: 16px;
    height: 16px;
    pointer-events: none;
}

.search-input {
    width: 100%;
    border-radius: 20px;
    height: 40px;
    border: 1px solid #00000080;
    padding: 0 16px 0 36px;
    background-color: #F1F4FF;
    color: #333;
    font-size: 14px;
    outline: none;
}

.search-input::placeholder {
    color: #BBBBBB;
}

.header-right {
    display: flex;
    align-items: center;
    gap: 20px;
    flex-shrink: 0;
}

.header-nav {
    text-decoration: none;
    font-family: "Manrope";
    font-weight: 400;
    font-size: clamp(14px, 1.5vw, 16px);
    line-height: 140%;
    color: #00000080;
    white-space: nowrap;
}

.header-nav:hover {
    color: #3383FB;
}

.notification-btn {
    background: none;
    border: 1px solid #BBBBBB;
    border-radius: 50%;
    width: 32px;
    height: 32px;
    cursor: pointer;
    font-weight: bold;
    flex-shrink: 0;
}

.avatar {
    width: 48px;
    height: 48px;
    border-radius: 50%;
    object-fit: cover;
    flex-shrink: 0;
}

/* Адаптивность */
@media (max-width: 768px) {
    .header-content {
        flex-wrap: wrap;
    }
    
    .header-left {
        flex: 1 1 auto;
    }
    
    .header-right {
        gap: 10px;
    }
    
    .header-nav {
        font-size: 14px;
    }
    
    .avatar {
        width: 40px;
        height: 40px;
    }
    .search-wrapper {
        display: none;
    }
}

@media (max-width: 480px) {
    .header-nav {
        display: none;
    }
    
    .logo {
        font-size: 16px;
    }
}
</style>