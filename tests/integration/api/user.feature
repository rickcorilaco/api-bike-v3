Feature: User API

    Background:
        * url baseURL
        * def randomNumber = read('./../helper/cpf.js')

    @create_user
    Scenario: should register an user with sucess
        * def name = karate.get('name', 'John Silva')
        * def username = karate.get('username', 'jsilva_'+randomNumber())
        * def password = karate.get('password', randomNumber())

        * def user =
        """
            {
                name: '#(name)',
                username: '#(username)',
                password: '#(password)'
            }
        """

        Given path '/users'
        And request user
        When method POST
        Then status 201
        And match response contains { id: '#notnull', name: '#(user.name)', username: '#(user.username)' }
    
    Scenario: should not create an duplicate user
        * def username = 'rjunior_' + randomNumber()
        * def password = randomNumber()
        * def user = { name: 'Robert Junior', username: '#(username)', password: '#(password)'}
        * call read('user.feature@create_user') user
       
        Given path '/users'
        And request user
        When method POST
        Then status 409

    Scenario: should login an user with success
        * def username = 'msantos_' + randomNumber()
        * def password = randomNumber()
        
        * def user = { name: 'Mary Santos', username: '#(username)', password: '#(password)'}
        * call read('user.feature@create_user') user
        
        * def login = { username: '#(username)', password: '#(password)' }

        Given path '/users/login'
        And request login
        When method POST
        Then status 200
        And match response contains { token: '#notnull' }

    Scenario: should not login an non-existent user
        * def username = 'notfound_' + randomNumber()
        * def password = randomNumber()    
        * def login = { username: '#(username)', password: '#(password)' }

        Given path '/users/login'
        And request login
        When method POST
        Then status 401
    
    
