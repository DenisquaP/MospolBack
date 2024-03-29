CREATE TABLE IF NOT EXISTS authors (
    author_id serial primary key,
    email varchar(120) UNIQUE NOT NULL,
    author_name varchar(120) UNIQUE NOT NULL,
    password varchar(240) NOT NULL,
    is_moderator boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS articles (
    article_id serial primary key,
    header TEXT NOT NULL,
    content TEXT NOT NULL,
    author INT NOT NULL,
    FOREIGN KEY (author) REFERENCES authors (author_id)
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id serial PRIMARY KEY,
    comment TEXT NOT NULL,
    commentator INT NOT NULL,
    article INT NOT NULL,
    approved boolean default false NOT NULL,
    FOREIGN KEY (commentator) REFERENCES authors (author_id),
    FOREIGN KEY (article) REFERENCES articles (article_id)
);

INSERT INTO
    authors (email, author_name, password, is_moderator)
VALUES
    ('some_email@mail.ru', 'Denis', 'some_password', true),
    ('some_email2@mail.ru', 'Anna', 'very_strong_pass', true);

INSERT INTO
    articles (header, content, author)
VALUES
    ('About animals', 'Cats are cool', 1),
    (
        'About animals 2',
        'Dogs are loud and annoying',
        1
    );

INSERT INTO
    articles (header, content, author)
VALUES
    (
        'Введение в программирование',
        'Это первая статья в серии, посвященной основам программирования.',
        1
    ),
    (
        'Основы баз данных',
        'В этой статье мы рассмотрим основные понятия и принципы работы баз данных.',
        2
    ),
    (
        'Искусственный интеллект: прошлое, настоящее, будущее',
        'В данной статье мы рассмотрим историю развития искусственного интеллекта и его текущие достижения.',
        2
    ),
    (
        'Веб-разработка: от HTML до JavaScript',
        'В этой статье мы расскажем об основах веб-разработки и важных концепциях веб-технологий.',
        1
    ),
    (
        'Криптовалюты и блокчейн: введение',
        'В данной статье мы изучим основы криптовалют и технологии блокчейн.',
        2
    ),
    (
        'Машинное обучение: основы и применение',
        'В этой статье мы рассмотрим основные концепции машинного обучения и его применение в различных областях.',
        2
    ),
    (
        'Агил-методология: эффективное управление проектами',
        'В данной статье мы расскажем о принципах и преимуществах агил-методологии в управлении проектами.',
        1
    ),
    (
        'Безопасность в сети Интернет: защита от угроз',
        'В этой статье мы рассмотрим основные принципы обеспечения безопасности в сети Интернет и методы защиты от угроз.',
        2
    ),
    (
        'Разработка мобильных приложений: выбор платформы',
        'В данной статье мы рассмотрим различные платформы для разработки мобильных приложений и поможем выбрать наиболее подходящую.',
        1
    ),
    (
        'Основы алгоритмов и структур данных',
        'В этой статье мы изучим основные алгоритмы и структуры данных, используемые в программировании.',
        1
    ),
    (
        'Интернет вещей: связывая физический и цифровой мир',
        'В данной статье мы рассмотрим принципы и применение интернета вещей в различных сферах.',
        2
    ),
    (
        'Разработка игр: от идеи до реализации',
        'В этой статье мы расскажем о процессе разработки компьютерных игр и важных этапах этого процесса.',
        2
    ),
    (
        'Операционные системы: принципы работы и функциональность',
        'В данной статье мы изучим основные принципы работы операционных систем и их функциональные возможности.',
        1
    ),
    (
        'Технологии облачных вычислений: преимущества и решения',
        'В этой статье мы рассмотрим преимущества облачных вычислений и различные решения для работы в облаке.',
        2
    ),
    (
        'Разработка веб-сайтов: от макета до запуска',
        'В данной статье мы расскажем о процессе разработки веб-сайтов и важных этапах этого процесса.',
        1
    ),
    (
        'Большие данные и аналитика: преимущества и технологии',
        'В этой статье мы изучим преимущества использования больших данных и основные технологии аналитики данных.',
        1
    ),
    (
        'Интерфейсы пользователя: отличия и тренды',
        'В данной статье мы расскажем о различных типах пользовательских интерфейсов и текущих трендах в их разработке.',
        2
    ),
    (
        'Разработка приложений для мобильных устройств: выбор платформы',
        'В этой статье мы поможем определиться с выбором платформы для разработки мобильных приложений.',
        2
    ),
    (
        'Кибербезопасность: защита от угроз и уязвимостей',
        'В данной статье мы рассмотрим методы и технологии обеспечения кибербезопасности и защиты от угроз и уязвимостей.',
        1
    ),
    (
        'Разработка интеллектуальных систем: принципы и применение',
        'В этой статье мы расскажем о принципах разработки интеллектуальных систем и их применении в различных областях.',
        2
    );