/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/*.{html,go}',
  ],
  theme: {
    extend: {},
  },
  safelist: [
    'fixed', 'inset-0', 'flex', 'items-center', 'justify-center', 'bg-gray-500', 'bg-opacity-50', 'z-50', "text-xl", "font-semibold", "text-gray-800", "bg-white", "rounded-lg", "w-96", "p-6", "text-gray-600", "mt-2", "justify-end", "mt-4", "space-x-4", "px-4", "py-2", "bg-gray-300", "rounded-md", "hover:bg-gray-400", "bg-red-600", "text-white", "rounded-md", "hover:bg-red-700", "bg-green-600", "p-4", "text-center", "bg-red-500", , "p-1", "rounded", "hover:bg-red-600", "bg-blue-500", "hover:bg-blue-600", "top-8", "left-1/2", "transform", "-translate-x-1/2", "text-sm", "font-medium", "py-4", "px-6", "shadow-lg", "transition-opacity", "duration-500", "ease-in-out",
    "opacity-0", "opacity-100", 'absolute', "top-1/2", "-translate-y-1/2", 'z-10', "bg-gray-600", "bg-blue-600", "bg-blue-700", "hover:bg-blue-700", "min-w-[150px]", "ml-2", "mr-2", "w-14", "absolute left-1", "top-1", "w-5", "h-5", "rounded-full", "transition-transform", "duration-300", "relative", "h-7", "cursor-pointer", "transition", "bg-gray-400", "translate-x-7", "m-auto", "bg-red-500", "inline-block", "px-3"
  ],
  plugins: [],
}
