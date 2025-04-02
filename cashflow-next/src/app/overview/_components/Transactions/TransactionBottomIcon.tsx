import { PiArrowDown, PiSpinner } from "react-icons/pi";

type TransactionBottomIconProps = {
  inView: boolean;
  isLoading: boolean;
  isFetchingNextPage: boolean;
};

export default function TransactionBottomIcon({
  inView,
  isLoading,
  isFetchingNextPage,
}: TransactionBottomIconProps) {
  return !inView || isLoading || isFetchingNextPage ? (
    <div className="-translate-x-1/2 absolute bottom-4 left-1/2 transform rounded-full border border-base-300 bg-base-100 p-2 text-gray-500 [&>svg]:h-4 [&>svg]:w-4">
      {!inView && !isLoading && !isFetchingNextPage && (
        <PiArrowDown className="animate-bounce " />
      )}
      {(isLoading || isFetchingNextPage) && (
        <PiSpinner className="animate-spin text-cente text-sm" />
      )}
    </div>
  ) : (
    <></>
  );
}
