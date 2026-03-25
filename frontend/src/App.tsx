import { useEffect, useState } from 'react';
import './App.css';

type HealthResponse = {
  status: string;
  message: string;
};

type PlayType =
  | 'strikeout'
  | 'flyout'
  | 'groundout'
  | 'walk'
  | 'hitByPitch'
  | 'single'
  | 'double'
  | 'triple'
  | 'homerun'
  | 'error'
  | 'sacrificeBunt'
  | 'steal';

type InningHalf = 'top' | 'bottom';

type BaseDestination = 'out' | 'first' | 'second' | 'third' | 'home';

type RunnerState = {
  first: boolean;
  second: boolean;
  third: boolean;
};

type TeamState = {
  teamId: string;
  score: number;
};

type GameState = {
  inning: number;
  inningHalf: InningHalf;
  outs: number;
  firstBattingTeamId: string;
  secondBattingTeamId: string;
  teams: Record<string, TeamState>;
  runner: RunnerState;
};

type AdvancementForm = {
  batter: BaseDestination | '';
  fromFirst: BaseDestination | '';
  fromSecond: BaseDestination | '';
  fromThird: BaseDestination | '';
};

type PlayPayload = {
  Type: PlayType;
  Override?: {
    Batter?: BaseDestination;
    FromFirst?: BaseDestination;
    FromSecond?: BaseDestination;
    FromThird?: BaseDestination;
  };
};

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

const playTypeOptions: Array<{ value: PlayType; label: string; category: string }> = [
  { value: 'strikeout', label: '三振', category: 'アウト' },
  { value: 'flyout', label: 'フライアウト', category: 'アウト' },
  { value: 'groundout', label: 'ゴロアウト', category: 'アウト' },
  { value: 'walk', label: '四球', category: '出塁' },
  { value: 'hitByPitch', label: '死球', category: '出塁' },
  { value: 'single', label: '単打', category: '安打' },
  { value: 'double', label: '二塁打', category: '安打' },
  { value: 'triple', label: '三塁打', category: '安打' },
  { value: 'homerun', label: '本塁打', category: '安打' },
  { value: 'error', label: '失策', category: '特殊' },
  { value: 'sacrificeBunt', label: '送りバント', category: '特殊' },
  { value: 'steal', label: '盗塁', category: '特殊' },
];

const destinationOptions: Array<{ value: BaseDestination; label: string }> = [
  { value: 'out', label: 'アウト' },
  { value: 'first', label: '一塁' },
  { value: 'second', label: '二塁' },
  { value: 'third', label: '三塁' },
  { value: 'home', label: 'ホームイン' },
];

const initialGameState = (): GameState => ({
  inning: 1,
  inningHalf: 'top',
  outs: 0,
  firstBattingTeamId: 'team1',
  secondBattingTeamId: 'team2',
  teams: {
    team1: { teamId: 'team1', score: 0 },
    team2: { teamId: 'team2', score: 0 },
  },
  runner: {
    first: false,
    second: false,
    third: false,
  },
});

const initialOverrideForm = (): AdvancementForm => ({
  batter: '',
  fromFirst: '',
  fromSecond: '',
  fromThird: '',
});

function App() {
  const [healthStatus, setHealthStatus] = useState<'loading' | 'ok' | 'error'>('loading');
  const [healthMessage, setHealthMessage] = useState('Connecting to backend...');
  const [selectedPlayType, setSelectedPlayType] = useState<PlayType>('single');
  const [overrideEnabled, setOverrideEnabled] = useState(false);
  const [overrideForm, setOverrideForm] = useState<AdvancementForm>(initialOverrideForm);
  const [gameState, setGameState] = useState<GameState>(initialGameState);
  const [queuedPlays, setQueuedPlays] = useState<PlayPayload[]>([]);

  useEffect(() => {
    const fetchHealth = async (): Promise<void> => {
      try {
        const response = await fetch(`${apiBaseUrl}/health`);
        if (!response.ok) {
          throw new Error(`HTTP ${response.status}`);
        }

        const data: HealthResponse = await response.json();
        setHealthStatus(data.status === 'ok' ? 'ok' : 'error');
        setHealthMessage(data.message);
      } catch (error) {
        setHealthStatus('error');
        setHealthMessage(error instanceof Error ? error.message : 'Unknown error');
      }
    };

    void fetchHealth();
  }, []);

  const selectedPlay = playTypeOptions.find((option) => option.value === selectedPlayType);

  const currentPayload = buildPlayPayload(selectedPlayType, overrideEnabled, overrideForm);
  const payloadPreview = JSON.stringify(
    {
      GameState: toBackendGameState(gameState),
      Play: currentPayload,
    },
    null,
    2,
  );

  const handleRunnerToggle = (base: keyof RunnerState): void => {
    setGameState((current) => ({
      ...current,
      runner: {
        ...current.runner,
        [base]: !current.runner[base],
      },
    }));
  };

  const handleScoreChange = (teamId: string, score: number): void => {
    setGameState((current) => ({
      ...current,
      teams: {
        ...current.teams,
        [teamId]: {
          ...current.teams[teamId],
          score,
        },
      },
    }));
  };

  const handleOverrideChange = (field: keyof AdvancementForm, value: BaseDestination | ''): void => {
    setOverrideForm((current) => ({
      ...current,
      [field]: value,
    }));
  };

  const queueCurrentPlay = (): void => {
    setQueuedPlays((current) => [currentPayload, ...current].slice(0, 8));
  };

  const clearOverride = (): void => {
    setOverrideEnabled(false);
    setOverrideForm(initialOverrideForm());
  };

  const resetWorkspace = (): void => {
    setGameState(initialGameState());
    setQueuedPlays([]);
    setSelectedPlayType('single');
    setOverrideEnabled(false);
    setOverrideForm(initialOverrideForm());
  };

  return (
    <main className="app-shell">
      <section className="hero-panel">
        <div className="hero-copy">
          <p className="eyebrow">Baseball Score Console</p>
          <h1>今の Go バックエンドにそのまま寄せた入力画面</h1>
          <p className="hero-text">
            `GameState` と `Play` を軸に、プレー記録と `Override` 編集を先に固める構成です。
            HTTP API はまだ `health` のみなので、記録内容はローカルで組み立てて JSON で確認できるようにしています。
          </p>
        </div>

        <div className="hero-meta">
          <div className={`status-pill status-pill-${healthStatus}`}>
            <span className="status-dot" />
            <span>Backend {healthStatus === 'ok' ? 'connected' : healthStatus}</span>
          </div>
          <div className="meta-card">
            <span className="meta-label">Health</span>
            <strong>{healthMessage}</strong>
          </div>
          <div className="meta-card">
            <span className="meta-label">API Base URL</span>
            <code>{apiBaseUrl}</code>
          </div>
        </div>
      </section>

      <section className="workspace-grid">
        <article className="panel">
          <div className="panel-header">
            <div>
              <p className="panel-kicker">Scoreboard</p>
              <h2>試合状態</h2>
            </div>
            <button className="ghost-button" type="button" onClick={resetWorkspace}>
              初期化
            </button>
          </div>

          <div className="scoreboard">
            <div className="team-row">
              <div>
                <span className="team-badge">TOP</span>
                <strong>{gameState.firstBattingTeamId}</strong>
              </div>
              <span className="score-value">{gameState.teams[gameState.firstBattingTeamId].score}</span>
            </div>
            <div className="team-row">
              <div>
                <span className="team-badge alt">BOT</span>
                <strong>{gameState.secondBattingTeamId}</strong>
              </div>
              <span className="score-value">{gameState.teams[gameState.secondBattingTeamId].score}</span>
            </div>
          </div>

          <div className="state-controls">
            <label>
              <span>Inning</span>
              <input
                min={1}
                type="number"
                value={gameState.inning}
                onChange={(event) =>
                  setGameState((current) => ({
                    ...current,
                    inning: Number(event.target.value) || 1,
                  }))
                }
              />
            </label>
            <label>
              <span>Half</span>
              <select
                value={gameState.inningHalf}
                onChange={(event) =>
                  setGameState((current) => ({
                    ...current,
                    inningHalf: event.target.value as InningHalf,
                  }))
                }
              >
                <option value="top">top</option>
                <option value="bottom">bottom</option>
              </select>
            </label>
            <label>
              <span>Outs</span>
              <input
                max={2}
                min={0}
                type="number"
                value={gameState.outs}
                onChange={(event) =>
                  setGameState((current) => ({
                    ...current,
                    outs: Math.min(2, Math.max(0, Number(event.target.value) || 0)),
                  }))
                }
              />
            </label>
          </div>

          <div className="diamond-card">
            <div className="diamond">
              <button
                className={`base base-home ${gameState.inningHalf === 'top' ? 'is-active' : ''}`}
                type="button"
              >
                {gameState.inningHalf}
              </button>
              <button
                className={`base base-first ${gameState.runner.first ? 'is-occupied' : ''}`}
                type="button"
                onClick={() => handleRunnerToggle('first')}
              >
                1B
              </button>
              <button
                className={`base base-second ${gameState.runner.second ? 'is-occupied' : ''}`}
                type="button"
                onClick={() => handleRunnerToggle('second')}
              >
                2B
              </button>
              <button
                className={`base base-third ${gameState.runner.third ? 'is-occupied' : ''}`}
                type="button"
                onClick={() => handleRunnerToggle('third')}
              >
                3B
              </button>
            </div>
            <p className="helper-text">塁ボタンで走者状態を切り替えます。</p>
          </div>

          <div className="score-edit-grid">
            <label>
              <span>{gameState.firstBattingTeamId} score</span>
              <input
                min={0}
                type="number"
                value={gameState.teams[gameState.firstBattingTeamId].score}
                onChange={(event) =>
                  handleScoreChange(gameState.firstBattingTeamId, Math.max(0, Number(event.target.value) || 0))
                }
              />
            </label>
            <label>
              <span>{gameState.secondBattingTeamId} score</span>
              <input
                min={0}
                type="number"
                value={gameState.teams[gameState.secondBattingTeamId].score}
                onChange={(event) =>
                  handleScoreChange(gameState.secondBattingTeamId, Math.max(0, Number(event.target.value) || 0))
                }
              />
            </label>
          </div>
        </article>

        <article className="panel">
          <div className="panel-header">
            <div>
              <p className="panel-kicker">Play Input</p>
              <h2>プレー選択</h2>
            </div>
            {selectedPlay ? <span className="category-chip">{selectedPlay.category}</span> : null}
          </div>

          <div className="play-grid">
            {playTypeOptions.map((option) => (
              <button
                key={option.value}
                className={`play-card ${selectedPlayType === option.value ? 'is-selected' : ''}`}
                type="button"
                onClick={() => setSelectedPlayType(option.value)}
              >
                <span>{option.label}</span>
                <small>{option.value}</small>
              </button>
            ))}
          </div>

          <div className="override-toggle-row">
            <div>
              <p className="panel-kicker">Advancement Override</p>
              <p className="helper-text">
                バックエンドの `Play.Override` に合わせて、必要なときだけ進塁先を上書きします。
              </p>
            </div>
            <label className="switch">
              <input
                checked={overrideEnabled}
                type="checkbox"
                onChange={(event) => setOverrideEnabled(event.target.checked)}
              />
              <span>{overrideEnabled ? 'enabled' : 'disabled'}</span>
            </label>
          </div>

          <div className={`override-grid ${overrideEnabled ? '' : 'is-disabled'}`}>
            <DestinationField
              label="Batter"
              value={overrideForm.batter}
              onChange={(value) => handleOverrideChange('batter', value)}
            />
            <DestinationField
              label="FromFirst"
              value={overrideForm.fromFirst}
              onChange={(value) => handleOverrideChange('fromFirst', value)}
            />
            <DestinationField
              label="FromSecond"
              value={overrideForm.fromSecond}
              onChange={(value) => handleOverrideChange('fromSecond', value)}
            />
            <DestinationField
              label="FromThird"
              value={overrideForm.fromThird}
              onChange={(value) => handleOverrideChange('fromThird', value)}
            />
          </div>

          <div className="action-row">
            <button className="primary-button" type="button" onClick={queueCurrentPlay}>
              このプレーをキューへ追加
            </button>
            <button className="ghost-button" type="button" onClick={clearOverride}>
              Override をクリア
            </button>
          </div>

          <div className="fit-card">
            <p className="panel-kicker">Backend Fit</p>
            <ul>
              <li>プレー種別は Go 側の `PlayType` に一致</li>
              <li>`Override` のフィールド名は `Batter / FromFirst / FromSecond / FromThird` に一致</li>
              <li>HTTP API が増えたら、このまま送信ペイロードとして接続可能</li>
            </ul>
          </div>
        </article>

        <article className="panel panel-preview">
          <div className="panel-header">
            <div>
              <p className="panel-kicker">Payload Preview</p>
              <h2>送信イメージ</h2>
            </div>
            <span className="queue-count">{queuedPlays.length} queued</span>
          </div>

          <pre className="json-preview">
            <code>{payloadPreview}</code>
          </pre>

          <div className="queue-panel">
            <div className="panel-header compact">
              <div>
                <p className="panel-kicker">Queued Plays</p>
                <h2>記録候補</h2>
              </div>
            </div>

            {queuedPlays.length === 0 ? (
              <p className="empty-state">まだプレーは追加されていません。選択中のプレーをキューへ追加してください。</p>
            ) : (
              <ol className="queue-list">
                {queuedPlays.map((play, index) => (
                  <li key={`${play.Type}-${index}`}>
                    <strong>{play.Type}</strong>
                    <code>{JSON.stringify(play)}</code>
                  </li>
                ))}
              </ol>
            )}
          </div>
        </article>
      </section>
    </main>
  );
}

type DestinationFieldProps = {
  label: string;
  value: BaseDestination | '';
  onChange: (value: BaseDestination | '') => void;
};

function DestinationField({ label, value, onChange }: DestinationFieldProps) {
  return (
    <label>
      <span>{label}</span>
      <select value={value} onChange={(event) => onChange(event.target.value as BaseDestination | '')}>
        <option value="">default</option>
        {destinationOptions.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </label>
  );
}

function buildPlayPayload(
  selectedPlayType: PlayType,
  overrideEnabled: boolean,
  overrideForm: AdvancementForm,
): PlayPayload {
  const payload: PlayPayload = {
    Type: selectedPlayType,
  };

  if (!overrideEnabled) {
    return payload;
  }

  const override = {
    ...(overrideForm.batter ? { Batter: overrideForm.batter } : {}),
    ...(overrideForm.fromFirst ? { FromFirst: overrideForm.fromFirst } : {}),
    ...(overrideForm.fromSecond ? { FromSecond: overrideForm.fromSecond } : {}),
    ...(overrideForm.fromThird ? { FromThird: overrideForm.fromThird } : {}),
  };

  if (Object.keys(override).length > 0) {
    payload.Override = override;
  }

  return payload;
}

function toBackendGameState(gameState: GameState) {
  return {
    Inning: gameState.inning,
    InningHalf: gameState.inningHalf,
    Outs: gameState.outs,
    FirstBattingTeamID: gameState.firstBattingTeamId,
    SecondBattingTeamID: gameState.secondBattingTeamId,
    Teams: {
      [gameState.firstBattingTeamId]: {
        TeamID: gameState.teams[gameState.firstBattingTeamId].teamId,
        Score: gameState.teams[gameState.firstBattingTeamId].score,
      },
      [gameState.secondBattingTeamId]: {
        TeamID: gameState.teams[gameState.secondBattingTeamId].teamId,
        Score: gameState.teams[gameState.secondBattingTeamId].score,
      },
    },
    Runner: {
      First: gameState.runner.first,
      Second: gameState.runner.second,
      Third: gameState.runner.third,
    },
  };
}

export default App;
