<template>
  <!-- разметка -->
   <div class="main_wrapper">
   <div class="left_side">
        <!-- <div class="left_content"><h1>q</h1></div> -->
        <img src="@/assets/img/enter.png" alt="">
   </div>
   <div class="right_side">
        <div class="form-wrapper">

            <form class="form" @submit.prevent="handleSubmit">

            <h2 class="form-title">Регистрация</h2>

            <!-- Форма собственности -->
            <div class="form-group">
                <div class="label-row">
                <label>Укажите вашу форму собственности *</label>

                <div class="hint">
                    <span>Для чего это нужно?</span>
                    <div class="tooltip">?</div>
                </div>
                </div>

                <div class="options two">
                <button 
                    type="button" 
                    class="option"
                    :class="{ active: ownership === 'ИП' }"
                    @click="ownership = 'ИП'"
                >
                    ИП
                </button>
                <button 
                    type="button" 
                    class="option"
                    :class="{ active: ownership === 'ООО' }"
                    @click="ownership = 'ООО'"
                >
                    ООО
                </button>
                </div>
            </div>

            <!-- Налогообложение -->
            <div class="form-group">
                <label>Выберите систему налогообложения *</label>

                <div class="options four">
                <button
                    type="button"
                    class="option"
                    :class="{ active: tax === 'УСН 6%' }"
                    @click="tax = 'УСН 6%'"
                >
                    УСН 6%
                </button>

                <button
                    type="button"
                    class="option"
                    :class="{ active: tax === 'УСН 15%' }"
                    @click="tax = 'УСН 15%'"
                >
                    УСН 15%
                </button>

                <button
                    type="button"
                    class="option"
                    :class="{ active: tax === 'Патент' }"
                    @click="tax = 'Патент'"
                >
                    Патент
                </button>

                <button
                    type="button"
                    class="option"
                    :class="{ active: tax === 'ОСНО' }"
                    @click="tax = 'ОСНО'"
                >
                    ОСНО
                </button>
                </div>
            </div>

            <!-- Сотрудники -->
            <div class="form-group">
                <label>Сколько у вас сотрудников? *</label>

                <div class="options three">
                <button
                    type="button"
                    class="option"
                    :class="{ active: employees === 'Нет' }"
                    @click="employees = 'Нет'"
                >
                    Нет
                </button>

                <button
                    type="button"
                    class="option"
                    :class="{ active: employees === 'До 10' }"
                    @click="employees = 'До 10'"
                >
                    До 10
                </button>

                <button
                    type="button"
                    class="option"
                    :class="{ active: employees === 'Более 10' }"
                    @click="employees = 'Более 10'"
                >
                    Более 10
                </button>
                </div>
            </div>

            <!-- Сфера деятельности -->
            <div class="form-group">
                <label>Укажите сферу вашей деятельности *</label>

                <select class="select" v-model="activity" required>
                <option disabled>Выберите деятельность</option>
                <option value="IT">IT</option>
                <option value="Маркетинг">Маркетинг</option>
                </select>
            </div>

            <!-- Имя + Email -->
            <div class="form-row">
                <div class="input-group">
                <label>Как вас зовут? *</label>
                <input placeholder="Имя" v-model="name" />
                </div>

                <div class="input-group">
                <label>Адрес электронной почты *</label>
                <input placeholder="E-mail" v-model="email" required />
                </div>
            </div>
            <div class="form-row">
                <div class="input-group">
                <label>Пароль *</label>
                <input placeholder="Пароль" type="password" v-model="password" required />
                </div>
            </div>

            <!-- Кнопка -->
            <button type="submit" class="submit">
                Готово
            </button>

            <p class="policy">
                Нажимая на кнопку, я соглашаюсь с политикой обработки персональных данных.
            </p>

            <p class="required">
                * Поля обязательны для заполнения
            </p>

            </form>
        </div>
   </div>
   </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// выборы
const ownership = ref('ИП')
const tax = ref('УСН 6%')
const employees = ref('Нет')

// поля
const activity = ref('')
const name = ref('')
const email = ref('')
const password = ref('')

// отправка формы
const handleSubmit = async () => {
  // собираем данные
  const formData = {
    business_form: ownership.value,
    nalog_system: tax.value,
    employees: employees.value,
    business_sphere: activity.value,
    name: name.value,
    email: email.value,
    password: password.value
  }

  console.log('Отправляем:', formData)

  try {
    const response = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
    })

    if (!response.ok) {
      throw new Error('Ошибка отправки')
    }

    const result = await response.json()
    console.log('Ответ сервера:', result)

    // ✅ Редирект на страницу авторизации
    router.push('/login')

  } catch (error) {
    console.error('Ошибка:', error)
  }
}
</script>

<style>
/* стили */
* {
    font-family: 'Manrope', sans-serif;
    font-weight: 600;
}
body {
    margin: 0;
    padding: 0;
}
.option {
  transition: background-color 0.25s ease, color 0.25s ease;
}
.main_wrapper {
    display: flex;
}
.left_content {
    background: url("@/assets/img/enter.png");
}
.left_side {
    width: 30%;
}
.right_side {
    width: 70%;
}
.form-wrapper {
  display: flex;
  justify-content: center;
  padding: 110px 20px;
  font-family: 'Manrope', sans-serif;
}

.form {
  width: 100%;
  max-width: 580px;
}

.form-title {
  font-size: 20px;
  font-weight: 600;
  line-height: 140%;
  color: rgba(0,0,0,0.7);
  margin-bottom: 30px;
}

/* Группа */
.form-group {
  margin-bottom: 28px;
}

label {
  font-size: 12px;
  color: rgba(0,0,0,0.4);
  margin-bottom: 10px;
  display: block;
}

/* label + hint */
.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.hint {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: rgba(0,0,0,0.4);
}

.tooltip {
  width: 10px;
  height: 10px;
  background: #8E8E8E;
  border-radius: 50%;
  font-size: 7px;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Кнопки вариантов */
.options {
  display: grid;
  gap: 12px;
}

.options.two {
  grid-template-columns: 1fr 1fr;
}

.options.four {
  grid-template-columns: repeat(4, 1fr);
}

.options.three {
  grid-template-columns: repeat(3, 1fr);
}

.option {
  height: 39px;
  border-radius: 6px;
  background: #F1F4FF;
  border: none;
  cursor: pointer;
  font-size: 14px;
  transition: 0.2s;
}

.option.active {
  background: #3383FB;
  color: white;
}

/* Select */
.select {
  width: 100%;
  height: 40px;
  background: #F1F4FF;
  border: none;
  border-radius: 6px;
  padding: 0 12px;
  font-size: 12px;
  color: #818181;
  appearance: none;
}

/* Row */
.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 30px;
}

.input-group input {
  width: 100%;
  height: 40px;
  background: #F1F4FF;
  border: none;
  border-radius: 6px;
  padding: 0 12px;
  font-size: 12px;
  color: #818181;
}

/* Submit */
.submit {
  width: 100%;
  height: 40px;
  background: #3383FB;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  color: white;
  cursor: pointer;
  margin-bottom: 20px;
}

/* Текст под кнопкой */
.policy,
.required {
  font-size: 12px;
  color: rgba(0,0,0,0.4);
  text-align: center;
  margin-bottom: 6px;
}

/* Адаптив */
@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .options.four {
    grid-template-columns: repeat(2, 1fr);
  }

  .options.three {
    grid-template-columns: 1fr;
  }
}
</style>