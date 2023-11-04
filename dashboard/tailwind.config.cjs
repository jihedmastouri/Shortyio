/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./index.html", "./src/**/*.{ts,tsx}"],
	darkMode: "class",
	// important: true,
	theme: {
		extend: {
			colors: {
				"s-gopher": "#39bebc",
				"s-gopher-l": "rgb(41,190,176,0.2)",
				"s-grass": "#559d94",
				"s-grass-l": "rgb(85,157,148,0.2)",
				"s-light": "#eefcfa",
				"s-biege": "#f3eada",
			},
		},
	},
	plugins: [],
};
