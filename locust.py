from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def ws_test01(self):
       self.client.get('/ws',headers={'name':'YangZhengHai'})
    @task
    def ws_test02(self):
       self.client.get('/ws',headers={'name':'SmallBai'})