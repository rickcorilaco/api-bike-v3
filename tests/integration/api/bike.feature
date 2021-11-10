Feature: bike api

    Background:
        * url baseURL

    Scenario: should create a bike with success
        * def bike = { brand: 'Caloi', model: '10' }
        
        Given path '/bikes'
        And request bike
        When method POST
        Then status 201
        And match response contains { id : '#notnull', brand:  '#(bike.brand)', model: '#(bike.model)' }

    Scenario: should not create a bike without required fields
        * def bike = { brand: '', model: '' }
        
        Given path '/bikes'
        And request bike
        When method POST
        Then status 400