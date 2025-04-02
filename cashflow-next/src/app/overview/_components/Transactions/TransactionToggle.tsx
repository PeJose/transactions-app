import type { TransactionTarget } from "@/types";

type TransactionToggleProps = {
  target: TransactionTarget;
  setTarget: React.Dispatch<React.SetStateAction<TransactionTarget>>;
};

export default function TransactionToggle({
  target,
  setTarget,
}: TransactionToggleProps) {
  return (
    <div className="flex justify-start p-4">
      <label className="flex items-center space-x-2">
        <span>Outgoing</span>
        <input
          type="checkbox"
          className="toggle toggle-primary"
          checked={target === "to"}
          onChange={() => setTarget(target === "from" ? "to" : "from")}
        />
        <span>Incoming</span>
      </label>
    </div>
  );
}
