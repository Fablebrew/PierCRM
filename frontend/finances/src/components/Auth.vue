<template>
  <div class="content">

    <!-- Лого слева -->
    <div class="brand">
      Vedi Project
    </div>

    <!-- Карточка -->
    <div class="login-card">

      <h2 class="title">Войти</h2>

      <!-- Соц кнопки -->
      <div class="social-row">
        <button class="social-btn">
          <span class="icon apple"></span>
          <span>Apple</span>
        </button>

        <button class="social-btn">
          <span class="icon google"></span>
          <span>Google</span>
        </button>
      </div>

      <!-- Разделитель -->
      <div class="divider">
        <div class="line"></div>
        <span>или продолжить с почтой</span>
        <div class="line"></div>
      </div>

      <!-- Форма -->
      <form class="form" @submit.prevent="handleSubmit">
        <input
          type="email"
          placeholder="Email"
          class="input"
          v-model="email"
          required
        />

        <input
          type="password"
          placeholder="Пароль"
          class="input"
          v-model="password"
          required
        />

        <div class="forgot">
          <a href="/resetpassword" class="signUp-link">Забыли пароль?</a> 
        </div>

        <button 
            type="submit" 
            class="submit-btn"
            :disabled="loading"
            >
            <span v-if="!loading">Войти</span>
            <span v-else class="loader"></span>
        </button>

        <p class="bottom-text">
          Нет аккаунта? <a href="/register" class="signUp-link">Зарегистрируйтесь</a>
        </p>
      </form>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)

const email = ref('')
const password = ref('')

const handleSubmit = async () => {
    if (loading.value) return

    loading.value = true // включаем загрузку
    const formData = {
        email: email.value,
        password: password.value
    }
    console.log('Отправляем:', formData)
    try {
        const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData),
        // credentials: 'include',
        })

        if (!response.ok) {
            throw new Error('Ошибка отправки')
        }

        const result = await response.json()
        console.log('Ответ сервера:', result)
        // ✅ сохраняем JWT
        localStorage.setItem('token', result.token)

        // ✅ Редирект на страницу авторизации
        router.push('/workspace')

    } catch (error) {
        console.error('Ошибка:', error)
    } finally {
        loading.value = false // выключаем в любом случае
    }
}
</script>

<style scoped>
* {
    font-family: 'Manrope', sans-serif;
    font-weight: 600;
}
.signUp-link {
    text-decoration: none;
    color: #3383FB;
}
.content {
  position: relative;
  width: 100%;
  min-height: 100vh;
  background: url("@/assets/img/authBG.jpg") center/cover no-repeat;
  display: flex;
  align-items: center;
  justify-content: flex-end; /* смещаем карточку вправо */
  padding: 0 8%;
  box-sizing: border-box;
}

/* Лого */
.brand {
  position: absolute;
  left: 8%;
  font-family: 'Manrope', sans-serif;
  font-weight: 600;
  font-size: clamp(28px, 4vw, 44px);
  color: #fff;
}

/* Карточка */
.login-card {
  width: 100%;
  max-width: 464px;
  padding: 40px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  display: flex;
  flex-direction: column;
  gap: 28px;
  box-shadow: 0 20px 40px rgba(0,0,0,0.1);
}

/* Заголовок */
.title {
  font-family: 'Inter', sans-serif;
  font-weight: 600;
  font-size: 24px;
  text-align: center;
}

/* Соц кнопки */
.social-row {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.social-btn {
  flex: 1;
  min-width: 140px;
  height: 40px;
  background: rgba(0,0,0,0.05);
  border: none;
  border-radius: 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  font-size: 13px;
  transition: 0.2s;
}

.social-btn:hover {
  background: rgba(0,0,0,0.08);
}

.icon {
  width: 16px;
  height: 16px;
  border-radius: 4px;
}

.apple {
  background: black;
}

.google {
  background: linear-gradient(45deg, #F25022, #34A853);
}

/* Разделитель */
.divider {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 12px;
  color: rgba(0,0,0,0.4);
}

.line {
  flex: 1;
  height: 1px;
  background: rgba(0,0,0,0.1);
}

/* Форма */
.form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* Инпут */
.input {
  width: 100%;
  height: 40px;
  padding: 0 12px;
  background: rgba(0,0,0,0.05);
  border: none;
  border-radius: 12px;
  font-size: 14px;
  box-sizing: border-box;
}

.input:focus {
  outline: 1px solid #3383FB;
}

/* Забыли пароль */
.forgot {
  font-size: 14px;
  color: #3383FB;
  cursor: pointer;
  align-self: flex-start;
}

/* Кнопка */
.submit-btn {
  width: 100%;
  height: 40px;
  background: black;
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 16px;
  cursor: pointer;
  transition: 0.2s;
}

.submit-btn:hover {
  opacity: 0.9;
}

/* Текст снизу */
.bottom-text {
  font-size: 14px;
  text-align: center;
  color: rgba(0,0,0,0.4);
}

.submit-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loader {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top: 2px solid #ffffff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* 📱 Адаптив */
@media (max-width: 1024px) {
  .content {
    justify-content: center;
  }

  .brand {
    position: static;
    margin-bottom: 40px;
  }

  .login-card {
    max-width: 420px;
  }
}

@media (max-width: 600px) {
  .login-card {
    padding: 28px;
    border-radius: 18px;
  }

  .social-row {
    flex-direction: column;
  }
}
</style>