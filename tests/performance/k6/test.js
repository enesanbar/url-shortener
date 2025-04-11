import { sleep } from "k6";
import http from "k6/http";

let thresholds = {}
thresholds[`http_req_duration{url:${__ENV.HOSTNAME}/url-shortener/api/mappings/ZSDASZX}`] = [
    "p(95)<=50",
]
thresholds[`http_req_duration{url:${__ENV.HOSTNAME}/url-shortener/redirect/ZSDASZX}`] = [
    "p(95)<=50",
]

export const options = {
    discardResponseBodies: true,
    scenarios: {
        contacts: {
            executor: 'constant-arrival-rate',
            rate: 10, // 10 RPS
            duration: '120s',
            preAllocatedVUs: 50,
            maxVUs: 2000,
        },
    },

    thresholds: thresholds,
};

export function setup() {
    const url = `${__ENV.HOSTNAME}/url-shortener/api/mappings`;
    const payload = JSON.stringify({
        code: "ZSDASZX",
        url: 'https://www.example.com',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}

export default function main() {
    let response;

    // get url mapping
    response = http.get(
        `${__ENV.HOSTNAME}/url-shortener/api/mappings/ZSDASZX`
    );

    response = http.get(
        `${__ENV.HOSTNAME}/url-shortener/redirect/ZSDASZX`,
        {
            redirects: 0
        }
    );


}
