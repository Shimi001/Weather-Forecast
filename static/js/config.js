export const cityName = "Tokyo"; // Default city name

export let forecast = [];
export let currentIndex = 0;

export const date = new Date();
export let dayOfWeek = date.getDay(); // Get the current day of the week (0-6)
export const today = dayOfWeek;

// Colors for circle backgrounds
export const dayColors = {
    0: '#34495e',  // Sunday - Dark Blue
    1: '#e74c3c',  // Monday - Red
    2: '#9b59b6',  // Tuesday - Purple
    3: '#2ecc71',  // Wednesday - Green
    4: '#f1c40f',  // Thursday - Yellow
    5: '#e67e22',  // Friday - Orange
    6: '#3498db'   // Saturday - Light Blue
};

// Japanese days of the week
export const japaneseDays = {
    0: '日', // Sunday
    1: '月', // Monday
    2: '火', // Tuesday
    3: '水', // Wednesday
    4: '木', // Thursday
    5: '金', // Friday
    6: '土'  // Saturday
};
