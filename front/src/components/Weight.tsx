import { useEffect, useRef, useState } from "react";

const Weight = () => {
  const [data, setData] = useState<Response | null>(null);
  const [basket, setBasket] = useState<
    { fruit: string; price: number; quantity: number }[]
  >([]);
  const [loading, setLoading] = useState(false);

  // useEffect(() => {
  //   fetch("http://localhost:8090/test")
  //     .then((response) => response.json())
  //     .then((data) => setData(data));
  // }, []);

  useEffect(() => {
    const container = document.querySelector(".container");
    if (container) {
    }
  }, []);

  const handleDrop = (fruit: string) => {
    console.log(`${fruit} dropped on target`);

    (async () => {
      setLoading(true);
      const response = await fetch("http://localhost:8090/added?file=" + fruit);

      const data = await response.json();
      if (data.Error !== undefined) {
        alert(data.Error);
      } else {
        setBasket((prev) => {
          const index = prev.findIndex((i) => i.fruit === data.fruit);
          if (index !== -1) {
            prev[index].quantity += 1;
            return [...prev];
          } else {
            return [
              ...prev,
              { fruit: data.fruit, quantity: 1, price: data.price },
            ];
          }
        });
      }
      setLoading(false);
    })();
  };

  return (
    <>
      <div
        className={
          "absolute top-32 w-full text-center hidden opacity-0 transition-opacity " +
          (loading ? "!block opacity-100" : "")
        }
      >
        <div className="m-auto inline-block p-3 bg-gray-200 rounded-md text-black">
          Trwa analiza...
        </div>
      </div>
      <div className=" border-[0px] border-neutral-200 mb-2 rounded-md p-2 ">
        <div className="pl-10 w-full flex">
          <div className="pr-5">
            <WeightTarget onDrop={handleDrop} />
          </div>
          <div className="">
            <div className="w-[670px] flex flex-wrap align-middle items-center  ">
              {[
                "banana.png",
                "chery.png",
                "orange.png",
                "orange2.png",
                "pinaple.png",
              ].map((fruit) => (
                <div
                  className="w-[200px] children:h-[200px] "
                  key={fruit}
                  data-swapy-slot={fruit + "slot"}
                >
                  <Fruit file={fruit} />
                </div>
              ))}
            </div>
          </div>
          <div className="grow">
            <div>Lista zakupów</div>
            <table className="w-full">
              <tbody>
                {basket.map((item) => {
                  return (
                    <tr>
                      <td>{item.fruit}</td>
                      <td>{item.quantity} szt.</td>
                      <td>{item.price * item.quantity} PLN</td>
                    </tr>
                  );
                })}
                <tr>
                  <td>Suma:</td>
                  <td></td>
                  <td>
                    {basket
                      .reduce((p, c) => c.price * c.quantity + p, 0)
                      .toFixed(2)}{" "}
                    PLN
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </>
  );
};

export { Weight };

const Fruit = ({ file }: { file: string }) => {
  const fruitRef = useRef<HTMLDivElement>(null);
  const [offset, setOffset] = useState({ x: 0, y: 0 });
  const [startPosition, setStartPosition] = useState({ left: 0, top: 0 });

  useEffect(() => {
    const element = fruitRef.current;
    if (element) {
      const rect = element.getBoundingClientRect();
      setStartPosition({ left: rect.left, top: rect.top });
    }
  }, []);

  const handleDragStart = (e: React.DragEvent) => {
    const element = fruitRef.current;
    if (element) {
      const rect = element.getBoundingClientRect();
      setOffset({ x: e.clientX - rect.left, y: e.clientY - rect.top });

      element.style.position = "fixed";
      element.style.left = `${rect.left}px`;
      element.style.top = `${rect.top}px`;
      element.style.zIndex = "1000";
      element.style.width = `${rect.width}px`;
      element.style.height = `${rect.height}px`;
    }
  };

  const handleDrag = (e: React.DragEvent) => {
    const element = fruitRef.current;
    if (element) {
      element.style.left = `${e.clientX - offset.x}px`;
      element.style.top = `${e.clientY - offset.y}px`;
    }
  };

  const handleDragEnd = (e: React.DragEvent) => {
    e.preventDefault();
    const element = fruitRef.current;
    const dropTarget = document.querySelector(".drop-target");
    const container = dropTarget;

    if (element && dropTarget && container) {
      const targetRect = dropTarget.getBoundingClientRect();
      const containerRect = container.getBoundingClientRect();
      const elementRect = element.getBoundingClientRect();
      const bottom10px = elementRect.bottom - 10;

      console.table(elementRect);

      if (
        bottom10px >= targetRect.top &&
        elementRect.left >= targetRect.left &&
        elementRect.right <= targetRect.right &&
        elementRect.bottom <= targetRect.bottom
      ) {
        //element.style.left = `${targetRect.left - containerRect.left}px`;
        //element.style.top = `${targetRect.top - containerRect.top}px`;
        //element.style.zIndex = "";
        dropTarget.dispatchEvent(
          new CustomEvent("fruitDropped", { detail: { file } })
        );
      } else {
        element.style.position = "fixed";
        element.style.left = `${startPosition.left}px`;
        element.style.top = `${startPosition.top}px`;
        element.style.zIndex = "";
      }
    }
  };

  return (
    <div
      ref={fruitRef}
      draggable
      onDragStart={handleDragStart}
      onDrag={handleDrag}
      onDragEnd={handleDragEnd}
      onDragExit={() => {
        console.log("drag exit");
      }}
      onDragOver={(e) => {
        e.preventDefault();
      }}
      className="border-[0px] border-neutral-200 rounded-md flex justify-center items-center"
    >
      <img src={"/public/fruits/" + file} className="max-h-full max-w-full" />
    </div>
  );
};

const WeightTarget = ({ onDrop }: { onDrop: (fruit: string) => void }) => {
  useEffect(() => {
    const dropTarget = document.querySelector(".drop-target");
    const handleFruitDrop = (e: Event) => {
      const customEvent = e as CustomEvent;
      onDrop(customEvent.detail.file);
    };
    dropTarget?.addEventListener("fruitDropped", handleFruitDrop);
    return () => {
      dropTarget?.removeEventListener("fruitDropped", handleFruitDrop);
    };
  }, [onDrop]);

  return (
    <div className="relative w-[300px] ">
      <div className=" drop-target relative top-0 left-0 w-[300px] h-[300px] border-[0px] border-red-400 rounded-md z-10"></div>
      <img
        src="/public/waga.png"
        className=" absolute max-w-[300px] h-[300px] top-[200px]"
      />
    </div>
  );
};
