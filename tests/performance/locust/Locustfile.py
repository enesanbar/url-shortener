import http
import json

from locust import TaskSet, task, constant
from locust.contrib.fasthttp import FastHttpUser


class URLShortenerTaskSet(TaskSet):

    def __init__(self, parent):
        super().__init__(parent)
        self.code = 'no_code'

    def on_start(self):
        resp = self.client.post('/url-shortener/api/mappings', name='Create Mapping', json={
            'url': 'https://a-very-long-url.com/with/a/lot/of/characters/and/some/other/characters/that/make/the/url/very-long',
        })
        json_resp = json.loads(resp.text)
        print(json_resp)
        self.code = json_resp['data']['code']

    @task
    def get_mapping(self):
        resp = self.client.get(f'/url-shortener/api/mappings/{self.code}', name='Get Mapping')
        if resp.status_code != http.HTTPStatus.OK:
            print(json.loads(resp.text)['error'])

    @task
    def redirect_mapping(self):
        resp = self.client.get(f'/url-shortener/redirect/{self.code}',
                               name='Redirect URL', allow_redirects=False)
        if resp.status_code != http.HTTPStatus.MOVED_PERMANENTLY:
            print(resp.text)

    def on_stop(self):
        self.client.delete(f'/url-shortener/api/mappings/{self.code}', name='Delete Mapping')


class User(FastHttpUser):
    host = "http://url-shortener:9090"
    tasks = [URLShortenerTaskSet]
    wait_time = constant(0)
