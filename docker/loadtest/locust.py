from locust import HttpUser, TaskSet, task

class RentalTasks(HttpUser):
    @task
    def c1(self):
        self.client.get('rentals?price_min=9000&price_max=75000')	
    @task
    def c2(self):
        self.client.get('rentals?limit=3&offset=6')
    @task
    def c3(self):
        self.client.get('rentals?ids=3,4,5')
    @task
    def c4(self):
        self.client.get('rentals?near=33.64,-117.93')     
     @task
    def c5(self):
        self.client.get('rentals?sort=price')
       @task
    def c6(self):
        self.client.get('rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price')
        
#http://localhost:8080/
#docker run -p 8089:8089 -v $PWD:/mnt/locust/ locustio/locust -f /mnt/locust/locust.py
