import { HEIGHT_OF_BATCH_STATUS_BUTTON } from "../BatchStatusButton/constants";

export const calculateRenderTarget = (batchListSize: number, allocatedAreaHeight: number) => {
  const heightOfOneBatchStatusButton = HEIGHT_OF_BATCH_STATUS_BUTTON + 4; // 4 is gap height

  if (heightOfOneBatchStatusButton * batchListSize <= allocatedAreaHeight) {
    return { hasMore: false };
  }

  let showCount = 0;
  let usedHeight = 0;
  while (usedHeight + heightOfOneBatchStatusButton <= allocatedAreaHeight) {
    usedHeight += HEIGHT_OF_BATCH_STATUS_BUTTON + 4;
    showCount++;
  }
  // Remove the last element for showing "more" button
  showCount--;

  const restCount = batchListSize - showCount;

  return { hasMore: true, showCount, restCount };
};
