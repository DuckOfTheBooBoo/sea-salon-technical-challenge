import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function hourMinuteParser(time: string): Date {
  const [hour, minute] = time.split(':')
  return new Date(0, 0, 0, parseInt(hour), parseInt(minute))
}