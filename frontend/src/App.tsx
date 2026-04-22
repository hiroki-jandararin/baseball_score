import type { ReactNode } from 'react';
import heroImage from './assets/hero.png';

const tabs = ['主要', '国内', '経済', 'IT', 'スポーツ', 'エンタメ'];

const featured = {
  title: '終盤の一打で接戦制す、Sharksが5-3で勝利',
  source: 'Team Sports',
  time: '18分前',
  image: heroImage,
  summary: 'SharksはRivalsとの接戦を5-3で制し、勝負どころの集中力と堅い守備で流れを渡さなかった。',
  lead: 'この試合では、1番山田が出塁で攻撃の起点を作り、4番佐藤がチャンスで2打点を挙げた。終盤まで拮抗した展開の中、守備陣も大きく崩れず、リードを守り切った。',
  paragraphs: [
    '試合は序盤から両チームが得点を重ねる展開となった。Sharksは山田が先頭で流れを作り、続く打線が走者を進めて得点機を広げた。',
    '勝負を分けたのは中盤以降の集中力だった。佐藤がチャンスで確実に打点を挙げ、相手に傾きかけた流れを引き戻した。',
    '守備面では大きなミスを抑え、終盤のピンチでも落ち着いてアウトを積み重ねた。攻守のバランスが勝利につながった一戦といえる。',
    'MVPには2打点の佐藤を選出。4番として結果を出し、チームに勝利を引き寄せる働きを見せた。',
  ],
};

type WithClassName = {
  className?: string;
  children: ReactNode;
};

type IconProps = {
  className?: string;
};

function IconBase({ className = '', children }: WithClassName) {
  return (
    <svg
      aria-hidden="true"
      className={className}
      fill="none"
      stroke="currentColor"
      strokeLinecap="round"
      strokeLinejoin="round"
      strokeWidth="2"
      viewBox="0 0 24 24"
    >
      {children}
    </svg>
  );
}

function Search({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <circle cx="11" cy="11" r="7" />
      <path d="m20 20-3.5-3.5" />
    </IconBase>
  );
}

function Bell({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <path d="M18 8a6 6 0 0 0-12 0c0 7-3 7-3 9h18c0-2-3-2-3-9" />
      <path d="M10 21h4" />
    </IconBase>
  );
}

function Home({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <path d="m3 11 9-8 9 8" />
      <path d="M5 10v10h14V10" />
      <path d="M10 20v-6h4v6" />
    </IconBase>
  );
}

function Newspaper({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <path d="M4 5h13a3 3 0 0 1 3 3v11H7a3 3 0 0 1-3-3Z" />
      <path d="M7 9h7" />
      <path d="M7 13h10" />
      <path d="M7 17h6" />
    </IconBase>
  );
}

function Bookmark({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <path d="M6 4h12v17l-6-4-6 4Z" />
    </IconBase>
  );
}

function User({ className }: IconProps) {
  return (
    <IconBase className={className}>
      <circle cx="12" cy="8" r="4" />
      <path d="M4 21a8 8 0 0 1 16 0" />
    </IconBase>
  );
}

function Card({ className = '', children }: WithClassName) {
  return <article className={`border bg-white ${className}`}>{children}</article>;
}

function CardContent({ className = '', children }: WithClassName) {
  return <div className={className}>{children}</div>;
}

function Badge({ className = '', children }: WithClassName) {
  return <span className={`inline-flex items-center text-xs font-bold ${className}`}>{children}</span>;
}

function Header() {
  return (
    <header className="sticky top-0 z-30 border-b border-neutral-200 bg-white/95 backdrop-blur">
      <div className="mx-auto flex max-w-md items-center justify-between px-4 py-3">
        <div>
          <div className="text-[11px] tracking-[0.18em] text-neutral-500">NEWS</div>
          <div className="text-xl font-extrabold tracking-tight text-[#06C755]">ニュース</div>
        </div>
        <div className="flex items-center gap-2">
          <button className="rounded-full p-2 hover:bg-neutral-100" type="button" aria-label="検索">
            <Search className="h-5 w-5 text-neutral-700" />
          </button>
          <button className="rounded-full p-2 hover:bg-neutral-100" type="button" aria-label="通知">
            <Bell className="h-5 w-5 text-neutral-700" />
          </button>
        </div>
      </div>
    </header>
  );
}

function Tabs() {
  return (
    <div className="sticky top-[73px] z-20 border-b border-neutral-200 bg-white/95 backdrop-blur">
      <div className="scrollbar-hide mx-auto flex max-w-md gap-5 overflow-x-auto whitespace-nowrap px-4 py-3 text-sm">
        {tabs.map((tab, index) => (
          <button
            key={tab}
            className={index === 4 ? 'font-bold text-neutral-950' : 'text-neutral-500'}
            type="button"
          >
            {tab}
          </button>
        ))}
      </div>
    </div>
  );
}

function FeaturedArticle() {
  return (
    <main className="px-4 pb-24 pt-4">
      <Card className="overflow-hidden rounded-3xl border-neutral-200 shadow-sm">
        <div className="relative">
          <img src={featured.image} alt={featured.title} className="h-60 w-full object-cover" />
          <Badge className="absolute left-3 top-3 rounded-full bg-[#06C755] px-3 py-1 text-white hover:bg-[#06C755]">
            注目
          </Badge>
        </div>

        <CardContent className="p-5">
          <div className="mb-3 flex items-center gap-2 text-xs text-neutral-400">
            <span>{featured.source}</span>
            <span>・</span>
            <span>{featured.time}</span>
          </div>

          <h1 className="text-[25px] font-bold leading-tight tracking-tight text-neutral-950">{featured.title}</h1>

          <p className="mt-4 text-[15px] font-medium leading-7 text-neutral-800">{featured.summary}</p>

          <p className="mt-4 border-l-4 border-[#06C755] pl-4 text-[15px] leading-7 text-neutral-700">
            {featured.lead}
          </p>

          <div className="mt-6 space-y-5">
            {featured.paragraphs.map((paragraph) => (
              <p key={paragraph} className="text-[15px] leading-8 text-neutral-700">
                {paragraph}
              </p>
            ))}
          </div>
        </CardContent>
      </Card>
    </main>
  );
}

function BottomNav() {
  const items = [
    { icon: Home, label: 'ホーム', active: false },
    { icon: Newspaper, label: 'ニュース', active: true },
    { icon: Bookmark, label: '保存', active: false },
    { icon: User, label: 'マイ', active: false },
  ];

  return (
    <nav className="sticky bottom-0 border-t border-neutral-200 bg-white/95 backdrop-blur">
      <div className="mx-auto grid max-w-md grid-cols-4 px-2 py-2">
        {items.map((item) => {
          const Icon = item.icon;
          return (
            <button key={item.label} className="flex flex-col items-center gap-1 rounded-2xl py-2 text-xs" type="button">
              <Icon className={`h-5 w-5 ${item.active ? 'text-[#06C755]' : 'text-neutral-500'}`} />
              <span className={item.active ? 'font-bold text-[#06C755]' : 'text-neutral-500'}>{item.label}</span>
            </button>
          );
        })}
      </div>
    </nav>
  );
}

export default function LineNewsStylePreview() {
  return (
    <div className="min-h-screen bg-[#f7f7f8]">
      <div className="mx-auto max-w-md bg-white shadow-[0_0_0_1px_rgba(0,0,0,0.03),0_12px_40px_rgba(0,0,0,0.06)]">
        <Header />
        <Tabs />
        <FeaturedArticle />
        <BottomNav />
      </div>
    </div>
  );
}
