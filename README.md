# BankGo

   RestApi Bank Server

  Возможности банка( В файлике RequestsForBank.postman_collection.json можно увидеть все виды http запросов)
  -----------


  /create
  ------------------

  Создает новый аккаунт пользователю. Если аккаунт уже есть, то выкидывает ошибку.

  /deposit
  -------------

 Начисляет деньги на аккаунт пользователя.

  /balance
  ------------

 Узнаёт баланс пользователя.

  /reserve
  ---------
  Резервирует на отдельном счету деньги. Если такой закказ уже был заразервирован или оплачен, то выбрасывает ошибку.
  
  /debit
  -----------------------------

  Списывает зарезервированные деньги
  
   Как устроена база данных?
  -----------------------------
  1. Transactions. Сюда записываются все транзакции, которые производятся со счётом пользователя. А именно, все покупки. Тут мы резервируем наш заказ, а потом если всё успешно, то списывает счёт. С помощью этой таблицы можно легко предоставлять отчёт бугалтерии(есть поле curtime)  а также делать unreserved.
  ![image](https://user-images.githubusercontent.com/79725120/200938045-186f8225-835f-4600-a99b-ebee9cf3c0f1.png)
  2. Balances. Тут хранятся все аккаунты пользователей и текущая сумма на них.
  ![image](https://user-images.githubusercontent.com/79725120/200938140-381f01bc-0b97-4d0b-aa89-4762026f37f6.png)

  3. Charges. Сюда мы записываем все зачисления на счёт пользователя.
  ![image](https://user-images.githubusercontent.com/79725120/200938236-29448d15-5f97-4f2b-a7f0-f7ac67814656.png)

 Запуск
 ----------------------------
 Для работы приложения нужно запустить 
 main.go а так же
 migrate файл, который сформирует баззу данных.
