import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
    stages: [
        { duration: '10s', target: 100 }, // Ramp-up from 0 to 20 VUs in 10 seconds
        { duration: '10s', target: 100 }, // Hold at 20 VUs for 10 seconds
        { duration: '10s', target: 0 },  // Ramp-down from 20 to 0 VUs in 10 seconds
    ],
};
export default function () {
    let url = 'http://localhost:3000';
    
    
    let response = http.get(url);
    
    // Checking the response code is 200 (OK)
    check(response, {
        'status is 200': (r) => r.status === 200, 
    });

    sleep(1)
}


//    |-------\
//   |          \
//  |             \
// |               \