import ReactPlayer from "react-player";
import { useState, useEffect } from "react";
import { useWindowSize } from "usehooks-ts";
import { useHotkeys } from "react-hotkeys-hook";

export default function App() {
  const windowSize = useWindowSize();
  const [url, setUrl] = useState("https://www.youtube.com/watch?v=ysz5S6PUM-U");
  useHotkeys("ctrl+l", () =>
    setUrl("https://www.youtube.com/watch?v=psuRGfAaju4")
  );

  return (
    <div className=" flex m-0 h-full w-full">
      <ReactPlayer
        url={url}
        controls={true}
        width={windowSize.width}
        height={windowSize.height}
        className="h-full w-full"
      />
    </div>
  );
}
