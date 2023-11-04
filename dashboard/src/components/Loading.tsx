import "./styles.css"

export default function Loading() {
  return (
    <div
      className="w-full h-screen z-30 overflow-hidden bg-gray-700 opacity-50 flex flex-col items-center justify-center"
    >
      <div className="loader ease-linear rounded-full border-4 border-t-4 border-gray-200 h-12 w-12 mb-4"></div>
      <h2 className="text-center text-white text-xl font-semibold">Loading...</h2>
      <p className="w-1/3 text-center text-white">
        This may take a few seconds.
      </p>
    </div>
  );
}
