document.addEventListener('DOMContentLoaded', function () {
    // Assuming you're using Chart.js for metric charts
    const ctx1 = document.getElementById('cpuChart').getContext('2d');
    const cpuChart = new Chart(ctx1, {
        type: 'line',
        data: {
            labels: ['Pod1', 'Pod2', 'Pod3'],  // Example labels
            datasets: [{
                label: 'CPU Usage',
                data: [30, 50, 70],  // Example data points
            }]
        }
    });

    const ctx2 = document.getElementById('memoryChart').getContext('2d');
    const memoryChart = new Chart(ctx2, {
        type: 'line',
        data: {
            labels: ['Pod1', 'Pod2', 'Pod3'],
            datasets: [{
                label: 'Memory Usage',
                data: [400, 700, 1000],
            }]
        }
    });
});
