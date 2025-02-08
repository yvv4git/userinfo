<template>
    <div class="page-main">
        <!-- Заголовок страницы -->
        <header class="header">
            <h1 class="header-title">{{ $t('welcome') }}</h1>
            <p class="header-subtitle">{{ $t('simpleSpa') }}</p>
        </header>

        <!-- Основной контент -->
        <main class="content">
            <!-- Карточки -->
            <div class="card">
                <h2 class="card-title">{{ $t('card1Title') }}</h2>
                <p class="card-text">{{ $t('card1Text') }}</p>
            </div>
            <div class="card">
                <h2 class="card-title">{{ $t('card2Title') }}</h2>
                <p class="card-text">{{ $t('card2Text') }}</p>
            </div>
            <div class="card">
                <h2 class="card-title">{{ $t('card3Title') }}</h2>
                <p class="card-text">{{ $t('card3Text') }}</p>
            </div>
        </main>

        <!-- Футер -->
        <footer class="footer">
            <p class="footer-text">&copy; {{ $t('footerAbout') }}</p>
        </footer>
    </div>
</template>

<script>
export default {
    name: 'PageMain',
};
</script>

<style scoped>
/* Основные стили страницы */
.page-main {
    font-family: Arial, sans-serif;
    background-color: #f5f5f5;
    color: #333;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
}

/* Стили для заголовка */
.header {
    text-align: center;
    margin-bottom: 40px;
}

.header-title {
    font-size: 2.5rem;
    color: #444;
}

.header-subtitle {
    font-size: 1.2rem;
    color: #666;
}

/* Стили для основного контента */
.content {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 20px;
    width: 100%;
    max-width: 1200px;
}

/* Стили для карточек */
.card {
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 20px;
    width: 300px;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.card:hover {
    transform: translateY(-5px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.card-title {
    font-size: 1.5rem;
    color: #444;
    margin-bottom: 10px;
}

.card-text {
    font-size: 1rem;
    color: #666;
}

/* Стили для футера */
.footer {
    margin-top: auto;
    text-align: center;
    padding: 20px;
    background-color: #e0e0e0;
    width: 100%;
}

.footer-text {
    font-size: 0.9rem;
    color: #555;
}
</style>