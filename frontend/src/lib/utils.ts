import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function hourMinuteParser(time: string): Date {
  const [hour, minute] = time.split(':')
  return new Date(0, 0, 0, parseInt(hour), parseInt(minute))
}

export function generateAvailableTimes(openTime: string, closeTime: string, serviceDuration: number): string[] {
  const times: string[] = [];

  // Helper function to parse HH:MM to a Date object
  function parseTime(time: string): Date {
      const [hours, minutes] = time.split(':').map(Number);
      const date = new Date();
      date.setHours(hours, minutes, 0, 0);
      return date;
  }

  // Helper function to format Date object to HH:MM string
  function formatTime(date: Date): string {
      const hours = date.getHours().toString().padStart(2, '0');
      const minutes = date.getMinutes().toString().padStart(2, '0');
      return `${hours}:${minutes}`;
  }

  // Parse open and close times
  const openDate = parseTime(openTime);
  const closeDate = parseTime(closeTime);
  
  // Generate available times
  let currentTime = new Date(openDate);

  while (currentTime < closeDate) {
      times.push(formatTime(currentTime));
      currentTime.setMinutes(currentTime.getMinutes() + serviceDuration);
      if (times.length > 1000) { // Safety check to prevent infinite loop
          console.error('Too many iterations, possible infinite loop');
          break;
      }
  }

  return times;
}

export function getHourMinute(dateTimeString: string): string {
  // Use regular expression to extract hour and minute (reliable for specific format)
  const match = dateTimeString.match(/^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})Z$/);

  if (match) {
    const hours = match[4].padStart(2, "0"); // Extract and format hours
    const minutes = match[5].padStart(2, "0"); // Extract and format minutes
    return `${hours}:${minutes}`;
  } else {
    // Handle invalid format (optional)
    console.error("Invalid datetime string format", dateTimeString);
    return ""; // Or throw an error
  }
}