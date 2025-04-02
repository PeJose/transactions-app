type TabNavigationProps = {
  activeTab: string;
  setActiveTab: (tab: string) => void;
};

export enum TABS {
  BALANCE = "balance",
  SWIFT = "swift",
  SEPA = "sepa",
  MAP = "map",
}

export default function TabNavigation({
  activeTab,
  setActiveTab,
}: TabNavigationProps) {
  return (
    <div className="tabs tabs-lift">
      <Tab
        activeTab={activeTab}
        tab={TABS.BALANCE}
        setActiveTab={setActiveTab}
        label="Balance Over Time"
      />
      <Tab
        activeTab={activeTab}
        tab={TABS.MAP}
        setActiveTab={setActiveTab}
        label="Transactions Map"
      />
      <Tab
        activeTab={activeTab}
        tab={TABS.SEPA}
        setActiveTab={setActiveTab}
        label="Sepa Transactions"
      />
      <Tab
        activeTab={activeTab}
        tab={TABS.SWIFT}
        setActiveTab={setActiveTab}
        label="Swift Transactions"
      />
    </div>
  );
}

function Tab({
  activeTab,
  tab,
  setActiveTab,
  label,
}: {
  activeTab: string;
  tab: string;
  setActiveTab: (tab: string) => void;
  label: string;
}) {
  return (
    <button
      type="button"
      className={`tab ${activeTab === tab ? "tab-active" : ""}`}
      onClick={() => setActiveTab(tab)}
    >
      {label}
    </button>
  );
}
