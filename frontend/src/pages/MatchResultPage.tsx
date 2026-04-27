import { useMemo, useState } from 'react';

type PlayerForm = {
  player_id: string;
  batting_order: string;
  position: string;
  hits: number | string;
  at_bats: number | string;
  rbi: number | string;
  runs: number | string;
  walks: number | string;
  strikeouts: number | string;
  errors: number | string;
  good_play: boolean;
  highlight_moment: boolean;
  memo: string;
};

type MatchForm = {
  team_id: string;
  opponent_name: string;
  match_date: string;
  location: string;
  is_win: 'win' | 'lose';
  team_score: string;
  opponent_score: string;
  note: string;
};

type NumericKey = 'hits' | 'at_bats' | 'rbi' | 'runs' | 'walks' | 'strikeouts' | 'errors';

const emptyPlayer: PlayerForm = {
  player_id: '',
  batting_order: '',
  position: '',
  hits: 0,
  at_bats: 0,
  rbi: 0,
  runs: 0,
  walks: 0,
  strikeouts: 0,
  errors: 0,
  good_play: false,
  highlight_moment: false,
  memo: '',
};

const numberFields: Array<[NumericKey, string]> = [
  ['hits', '安打'],
  ['at_bats', '打数'],
  ['rbi', '打点'],
  ['runs', '得点'],
  ['walks', '四球'],
  ['strikeouts', '三振'],
  ['errors', '失策'],
];

function IconBadge({ children, dark = false }: { children: string; dark?: boolean }) {
  return (
    <div
      className={`flex h-11 w-11 items-center justify-center rounded-2xl text-lg font-black ${
        dark ? 'bg-slate-950 text-white' : 'bg-red-600 text-white'
      }`}
    >
      {children}
    </div>
  );
}

function Field({
  label,
  required,
  children,
}: {
  label: string;
  required?: boolean;
  children: React.ReactNode;
}) {
  return (
    <label className="block">
      <div className="mb-1.5 flex items-center gap-1 text-sm font-semibold text-slate-700">
        {label}
        {required && <span className="text-red-600">*</span>}
      </div>
      {children}
    </label>
  );
}

function inputClass(extra = '') {
  return `w-full rounded-xl border border-slate-200 bg-white px-3.5 py-2.5 text-sm text-slate-900 outline-none transition focus:border-red-500 focus:ring-4 focus:ring-red-100 ${extra}`;
}

export default function BaseballMatchResultForm() {
  const [match, setMatch] = useState<MatchForm>({
    team_id: '',
    opponent_name: '',
    match_date: '',
    location: '',
    is_win: 'win',
    team_score: '',
    opponent_score: '',
    note: '',
  });

  const [players, setPlayers] = useState<PlayerForm[]>([{ ...emptyPlayer }]);

  const resultLabel = match.is_win === 'win' ? 'WIN' : 'LOSE';
  const scoreText = useMemo(() => {
    if (match.team_score === '' || match.opponent_score === '') return '- : -';
    return `${match.team_score} : ${match.opponent_score}`;
  }, [match.team_score, match.opponent_score]);

  const updateMatch = (key: keyof MatchForm, value: string) => {
    setMatch((prev) => ({ ...prev, [key]: value }));
  };

  const updatePlayer = <K extends keyof PlayerForm>(index: number, key: K, value: PlayerForm[K]) => {
    setPlayers((prev) => prev.map((player, i) => (i === index ? { ...player, [key]: value } : player)));
  };

  const addPlayer = () => {
    setPlayers((prev) => [...prev, { ...emptyPlayer }]);
  };

  const removePlayer = (index: number) => {
    setPlayers((prev) => prev.filter((_, i) => i !== index));
  };

  const handleSubmit = () => {
    // 現状のバックエンドでは保存API未接続のため、MVPはペイロード生成まで。
    const payload = {
      match: {
        ...match,
        is_win: match.is_win === 'win' ? 1 : 0,
        match_date: match.match_date ? new Date(`${match.match_date}T00:00:00+09:00`).toISOString() : '',
        team_id: Number(match.team_id || 0),
        team_score: Number(match.team_score || 0),
        opponent_score: Number(match.opponent_score || 0),
      },
      player_stats: players.map((player) => ({
        ...player,
        player_id: Number(player.player_id || 0),
        batting_order: player.batting_order ? Number(player.batting_order) : null,
        hits: Number(player.hits || 0),
        at_bats: Number(player.at_bats || 0),
        rbi: Number(player.rbi || 0),
        runs: Number(player.runs || 0),
        walks: Number(player.walks || 0),
        strikeouts: Number(player.strikeouts || 0),
        errors: Number(player.errors || 0),
        good_play: player.good_play ? 1 : 0,
        highlight_moment: player.highlight_moment ? 1 : 0,
      })),
    };

    console.log('submit payload', payload);
    alert('送信ペイロードをconsoleに出力しました');
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-white via-slate-50 to-red-50 p-6 text-slate-900">
      <div className="mx-auto max-w-6xl space-y-6">
        <header className="overflow-hidden rounded-[2rem] border border-red-100 bg-white shadow-xl shadow-red-100/40">
          <div className="relative p-8">
            <div className="absolute right-0 top-0 h-40 w-40 rounded-bl-full bg-red-600/10" />
            <div className="absolute bottom-0 right-24 h-24 w-24 rounded-full bg-red-500/10" />
            <div className="relative flex flex-col gap-6 md:flex-row md:items-end md:justify-between">
              <div>
                <div className="mb-4 inline-flex items-center gap-2 rounded-full bg-red-50 px-4 py-2 text-sm font-bold text-red-700">
                  <span>⚾</span> Match Result Entry
                </div>
                <h1 className="text-3xl font-black tracking-tight md:text-5xl">試合結果登録</h1>
                <p className="mt-3 max-w-2xl text-sm leading-7 text-slate-600">
                  試合単位の情報と、選手ごとの成績をまとめて登録するMVP用フォームです。
                </p>
              </div>

              <div className="rounded-3xl border border-slate-100 bg-slate-950 p-5 text-white shadow-lg">
                <div className="mb-2 flex items-center gap-2 text-sm text-slate-300">
                  <span className="text-red-400">★</span> Result Preview
                </div>
                <div className="flex items-end gap-4">
                  <div className="text-4xl font-black">{scoreText}</div>
                  <div
                    className={`rounded-full px-3 py-1 text-sm font-black ${
                      match.is_win === 'win' ? 'bg-red-600' : 'bg-slate-700'
                    }`}
                  >
                    {resultLabel}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </header>

        <div className="grid gap-6 lg:grid-cols-[1fr_340px]">
          <main className="space-y-6">
            <section className="rounded-[2rem] border border-slate-100 bg-white shadow-lg shadow-slate-200/60">
              <div className="p-6">
                <div className="mb-6 flex items-center gap-3">
                  <IconBadge>日</IconBadge>
                  <div>
                    <h2 className="text-xl font-black">試合情報</h2>
                    <p className="text-sm text-slate-500">matches テーブルに登録する情報</p>
                  </div>
                </div>

                <div className="grid gap-5 md:grid-cols-2">
                  <Field label="チームID" required>
                    <input
                      className={inputClass()}
                      value={match.team_id}
                      onChange={(e) => updateMatch('team_id', e.target.value)}
                      placeholder="例: 1"
                    />
                  </Field>
                  <Field label="対戦相手" required>
                    <input
                      className={inputClass()}
                      value={match.opponent_name}
                      onChange={(e) => updateMatch('opponent_name', e.target.value)}
                      placeholder="例: Red Hawks"
                    />
                  </Field>
                  <Field label="試合日" required>
                    <input
                      type="date"
                      className={inputClass()}
                      value={match.match_date}
                      onChange={(e) => updateMatch('match_date', e.target.value)}
                    />
                  </Field>
                  <Field label="場所">
                    <div className="relative">
                      <span className="absolute left-3 top-2.5 text-slate-400">📍</span>
                      <input
                        className={inputClass('pl-9')}
                        value={match.location}
                        onChange={(e) => updateMatch('location', e.target.value)}
                        placeholder="例: 多摩川グラウンド"
                      />
                    </div>
                  </Field>
                  <Field label="勝敗" required>
                    <div className="grid grid-cols-2 gap-3">
                      <button
                        type="button"
                        onClick={() => updateMatch('is_win', 'win')}
                        className={`rounded-xl border px-4 py-2.5 text-sm font-black transition ${
                          match.is_win === 'win'
                            ? 'border-red-600 bg-red-600 text-white'
                            : 'border-slate-200 bg-white text-slate-600'
                        }`}
                      >
                        勝利
                      </button>
                      <button
                        type="button"
                        onClick={() => updateMatch('is_win', 'lose')}
                        className={`rounded-xl border px-4 py-2.5 text-sm font-black transition ${
                          match.is_win === 'lose'
                            ? 'border-slate-900 bg-slate-900 text-white'
                            : 'border-slate-200 bg-white text-slate-600'
                        }`}
                      >
                        敗北
                      </button>
                    </div>
                  </Field>
                  <div className="grid grid-cols-2 gap-3">
                    <Field label="自チーム得点" required>
                      <input
                        type="number"
                        min="0"
                        className={inputClass()}
                        value={match.team_score}
                        onChange={(e) => updateMatch('team_score', e.target.value)}
                      />
                    </Field>
                    <Field label="相手得点" required>
                      <input
                        type="number"
                        min="0"
                        className={inputClass()}
                        value={match.opponent_score}
                        onChange={(e) => updateMatch('opponent_score', e.target.value)}
                      />
                    </Field>
                  </div>
                  <div className="md:col-span-2">
                    <Field label="試合メモ">
                      <textarea
                        className={inputClass('min-h-28 resize-none')}
                        value={match.note}
                        onChange={(e) => updateMatch('note', e.target.value)}
                        placeholder="試合全体の流れ、反省点、印象的な場面など"
                      />
                    </Field>
                  </div>
                </div>
              </div>
            </section>

            <section className="rounded-[2rem] border border-slate-100 bg-white shadow-lg shadow-slate-200/60">
              <div className="p-6">
                <div className="mb-6 flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
                  <div className="flex items-center gap-3">
                    <IconBadge dark>人</IconBadge>
                    <div>
                      <h2 className="text-xl font-black">選手成績</h2>
                      <p className="text-sm text-slate-500">player_match_stats を人数分登録</p>
                    </div>
                  </div>
                  <button
                    type="button"
                    onClick={addPlayer}
                    className="rounded-xl bg-red-600 px-4 py-2.5 text-sm font-bold text-white transition hover:bg-red-700"
                  >
                    <span className="mr-2">＋</span> 選手を追加
                  </button>
                </div>

                <div className="space-y-5">
                  {players.map((player, index) => (
                    <div key={index} className="rounded-3xl border border-slate-100 bg-slate-50/70 p-5">
                      <div className="mb-5 flex items-center justify-between gap-3">
                        <div className="flex items-center gap-3">
                          <div className="flex h-10 w-10 items-center justify-center rounded-2xl bg-white text-lg font-black text-red-600 shadow-sm">
                            {index + 1}
                          </div>
                          <div>
                            <h3 className="font-black">Player Stat</h3>
                            <p className="text-xs text-slate-500">選手ID・打順・守備位置・打撃成績</p>
                          </div>
                        </div>
                        {players.length > 1 && (
                          <button
                            type="button"
                            onClick={() => removePlayer(index)}
                            className="rounded-xl border border-red-100 bg-white p-2 text-red-600 transition hover:bg-red-50"
                          >
                            <span className="text-lg leading-none">×</span>
                          </button>
                        )}
                      </div>

                      <div className="grid gap-4 md:grid-cols-4">
                        <Field label="選手ID" required>
                          <input
                            className={inputClass()}
                            value={player.player_id}
                            onChange={(e) => updatePlayer(index, 'player_id', e.target.value)}
                            placeholder="例: 12"
                          />
                        </Field>
                        <Field label="打順">
                          <input
                            type="number"
                            min="1"
                            max="9"
                            className={inputClass()}
                            value={player.batting_order}
                            onChange={(e) => updatePlayer(index, 'batting_order', e.target.value)}
                            placeholder="1〜9"
                          />
                        </Field>
                        <Field label="守備位置">
                          <input
                            className={inputClass()}
                            value={player.position}
                            onChange={(e) => updatePlayer(index, 'position', e.target.value)}
                            placeholder="例: SS / CF / P"
                          />
                        </Field>
                        <div className="grid grid-cols-2 gap-3">
                          <label className="flex items-center gap-2 rounded-xl border border-slate-200 bg-white px-3 py-2.5 text-sm font-bold text-slate-700">
                            <input
                              type="checkbox"
                              checked={player.good_play}
                              onChange={(e) => updatePlayer(index, 'good_play', e.target.checked)}
                              className="h-4 w-4 accent-red-600"
                            />
                            好プレー
                          </label>
                          <label className="flex items-center gap-2 rounded-xl border border-slate-200 bg-white px-3 py-2.5 text-sm font-bold text-slate-700">
                            <input
                              type="checkbox"
                              checked={player.highlight_moment}
                              onChange={(e) => updatePlayer(index, 'highlight_moment', e.target.checked)}
                              className="h-4 w-4 accent-red-600"
                            />
                            見せ場
                          </label>
                        </div>
                      </div>

                      <div className="mt-4 grid grid-cols-2 gap-3 md:grid-cols-7">
                        {numberFields.map(([key, label]) => (
                          <Field key={key} label={label} required={key === 'hits' || key === 'at_bats'}>
                            <input
                              type="number"
                              min="0"
                              className={inputClass()}
                              value={player[key]}
                              onChange={(e) => updatePlayer(index, key, e.target.value)}
                            />
                          </Field>
                        ))}
                      </div>

                      <div className="mt-4">
                        <Field label="選手メモ">
                          <textarea
                            className={inputClass('min-h-20 resize-none')}
                            value={player.memo}
                            onChange={(e) => updatePlayer(index, 'memo', e.target.value)}
                            placeholder="今日の良かった点、改善点、印象的なプレーなど"
                          />
                        </Field>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            </section>
          </main>

          <aside className="space-y-6">
            <section className="sticky top-6 rounded-[2rem] border border-slate-100 bg-white shadow-xl shadow-red-100/50">
              <div className="p-6">
                <div className="mb-5 rounded-3xl bg-slate-950 p-5 text-white">
                  <div className="mb-2 text-xs font-bold uppercase tracking-[0.2em] text-red-300">Summary</div>
                  <div className="text-3xl font-black">{scoreText}</div>
                  <div className="mt-2 text-sm text-slate-300">vs {match.opponent_name || '対戦相手未入力'}</div>
                </div>

                <div className="space-y-3 text-sm">
                  <div className="flex justify-between border-b border-slate-100 pb-3">
                    <span className="text-slate-500">登録選手数</span>
                    <span className="font-black">{players.length} 人</span>
                  </div>
                  <div className="flex justify-between border-b border-slate-100 pb-3">
                    <span className="text-slate-500">好プレー</span>
                    <span className="font-black">{players.filter((p) => p.good_play).length} 件</span>
                  </div>
                  <div className="flex justify-between border-b border-slate-100 pb-3">
                    <span className="text-slate-500">ハイライト</span>
                    <span className="font-black">{players.filter((p) => p.highlight_moment).length} 件</span>
                  </div>
                </div>

                <button
                  type="button"
                  onClick={handleSubmit}
                  className="mt-6 h-12 w-full rounded-2xl bg-red-600 text-base font-black text-white shadow-lg shadow-red-200 transition hover:bg-red-700"
                >
                  <span className="mr-2">✓</span> 試合結果を登録
                </button>

                <p className="mt-4 text-xs leading-5 text-slate-500">
                  現状のバックエンドでは保存APIが未接続のため、ここでは送信ペイロードの確認まで行います。
                </p>
              </div>
            </section>
          </aside>
        </div>
      </div>
    </div>
  );
}
