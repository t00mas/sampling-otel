import http from 'k6/http';

const TOTAL_RPS = 100;
const DURATION = '2h';

function rate(percentage) {
  return TOTAL_RPS * percentage;
}

export const options = {
    scenarios: {
      normal: {
        exec: 'normal_endpoint',
        executor: 'constant-arrival-rate',
        duration: DURATION,
        preAllocatedVUs: 40,
        rate: rate(0.8)
      },
      long: {
          exec: 'long_endpoint',
          executor: 'constant-arrival-rate',
          duration: DURATION,
          preAllocatedVUs: 10,
          rate: rate(0.1)
      },
      error: {
        exec: 'error_endpoint',
        executor: 'constant-arrival-rate',
        duration: DURATION,
        preAllocatedVUs: 5,
        rate: rate(0.1),
      }
    }
  };

export function normal_endpoint() {
    http.get('http://entrypoint-service:8080');
}

export function long_endpoint() {
    http.get('http://entrypoint-service:8080/long');
}

export function error_endpoint() {
    http.get('http://entrypoint-service:8080/error');
}