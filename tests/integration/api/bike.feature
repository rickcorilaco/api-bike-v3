Feature: Bike API

    Background:
        * url baseURL
    
    @list_bike
    Scenario: should get a list of bikes
        * def result = call read('bike.feature@create_bike') { brand: 'Caloi', model: '10' }
        * def bike1 = result.response
        * def result = call read('bike.feature@create_bike') { brand: 'Absolute', model: 'Mia' }        
        * def bike2 = result.response

        Given path '/bikes'
        When method GET
        Then status 200
        And match response contains deep bike1
        And match response contains deep bike2

    Scenario: should get a filtered list of bikes
        * def result = call read('bike.feature@list_bike')

        Given path '/bikes'
        And param brand = result.bike1.brand 
        When method GET
        Then status 200
        And match response contains result.bike1
        And match response !contains result.bike2

        Given path '/bikes'
        And param model = result.bike1.model 
        When method GET
        Then status 200
        And match response contains result.bike1
        And match response !contains result.bike2

        Given path '/bikes'
        And param brand = 'Crazy'
        And param model = 'Frog' 
        When method GET
        Then status 200
        And match response !contains result.bike1
        And match response !contains result.bike2

    Scenario: should get a bike
        * def result = call read('bike.feature@create_bike') { brand: 'Soul', model: '3R5' }
        * def bikeId = result.response.id
        * def expected = { id: '#(bikeId)', brand: '#(result.bike.brand)', model: '#(result.bike.model)' }

        Given path '/bikes/' + bikeId
        When method GET
        Then status 200
        And match response == expected

    Scenario: should not get a non-existent bike
        Given path '/bikes/f2cc51e4-3734-4e7b-9c38-ad9764c2048a'
        When method GET
        Then status 404

    @create_bike
    Scenario: should create a bike with success
        * def brand = karate.get('brand', 'Caloi')
        * def model = karate.get('model', '10')
        * def bike = { brand: '#(brand)', model: '#(model)' }
        
        Given path '/bikes'
        And request bike
        When method POST
        Then status 201
        And match response contains { id: '#notnull', brand:  '#(bike.brand)', model: '#(bike.model)' }

    Scenario: should not create a bike without required fields
        * def bike = { brand: '', model: '' }
        
        Given path '/bikes'
        And request bike
        When method POST
        Then status 400

    Scenario: should delete a bike
        * def result = call read('bike.feature@create_bike')
        * def bikeId = result.response.id

        Given path '/bikes/' + bikeId
        When method DELETE
        Then status 204

    Scenario: should not delete a non-existent bike
        Given path '/bikes/f2cc51e4-3734-4e7b-9c38-ad9764c2048a'
        When method DELETE
        Then status 404