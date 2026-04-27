export default function App() {
  return (
    <div className="min-h-screen bg-slate-50 px-6 py-12">
      <div className="mx-auto max-w-3xl rounded-3xl border border-slate-200 bg-white p-8 shadow-sm">
        <h1 className="text-2xl font-black text-slate-900">Baseball Score App</h1>
        <p className="mt-2 text-sm text-slate-600">表示したいページを選択してください。</p>

        <div className="mt-6 grid gap-3 sm:grid-cols-2">
          <a
            href="/entry"
            className="inline-flex items-center justify-center rounded-xl bg-red-600 px-4 py-3 text-sm font-bold text-white transition hover:bg-red-700"
          >
            結果入力ページへ
          </a>
          <a
            href="/news"
            className="inline-flex items-center justify-center rounded-xl bg-[#06C755] px-4 py-3 text-sm font-bold text-white transition hover:bg-[#05af4b]"
          >
            ニュースページへ
          </a>
        </div>
      </div>
    </div>
  );
}
