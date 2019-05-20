import * as PIXI from 'pixi.js'


function PikachuInit(){
    var $_canvas = $('#canvas')

        console.log({width: $_canvas.innerWidth(), height: $_canvas.innerHeight()})
        let app = new PIXI.Application({width: $_canvas.innerWidth(), height: $_canvas.innerHeight()});
        $_canvas.append(app.view);

        let img = new PIXI.Sprite.from("/assets/images/pikachu.jpg");
        img.width = $_canvas.innerWidth();
        img.height = $_canvas.innerHeight();
        app.stage.addChild(img);

        var depthMap = new PIXI.Sprite.from("/assets/images/pikachu-map2.jpg");
        depthMap.width = $_canvas.innerWidth();
        depthMap.height = $_canvas.innerHeight();
        app.stage.addChild(depthMap);

        var displacementFilter = new PIXI.filters.DisplacementFilter(depthMap);
        app.stage.filters = [displacementFilter]

        $_canvas.mousemove(function(e) {
            displacementFilter.scale.x = ( $_canvas.innerWidth() / 2 - e.clientX) /21;
            displacementFilter.scale.y = ( $_canvas.innerHeight() / 2 - e.clientY) /21;
        })
}

function Cars(n){
    var $car = $('#funny-car')

    const app = new PIXI.Application({width: $car.innerWidth(), height: $car.innerHeight()});
    $car.append(app.view);

    const cars = [];

    const totalcars = n;

    for (let i = 0; i < totalcars; i++) {

        const car = PIXI.Sprite.from('/assets/images/funny-car.png');


        car.anchor.set(0.5);

        car.scale.set(0.8 + Math.random() * 0.3);


        car.x = Math.random() * $car.innerWidth();
        car.y = Math.random() * $car.innerHeight();

        car.tint = Math.random() * 0xFFFFFF;


        car.direction = Math.random() * Math.PI * 2;

        car.turningSpeed = Math.random() - 0.8;

        car.speed = 2 + Math.random() * 2;

        cars.push(car);

        app.stage.addChild(car);
    }

    // create a bounding box for the little cars
    const carBoundsPadding = 100;
    const carBounds = new PIXI.Rectangle(-carBoundsPadding,
        -carBoundsPadding,
        $car.innerWidth() + carBoundsPadding * 2,
        $car.innerHeight() + carBoundsPadding * 2);

    app.ticker.add(() => {
        for (let i = 0; i < cars.length; i++) {
            const car = cars[i];
            car.direction += car.turningSpeed * 0.01;
            car.x += Math.sin(car.direction) * car.speed;
            car.y += Math.cos(car.direction) * car.speed;
            car.rotation = -car.direction - Math.PI / 2;

            if (car.x < carBounds.x) {
                car.x += carBounds.width;
            } else if (car.x > carBounds.x + carBounds.width) {
                car.x -= carBounds.width;
            }

            if (car.y < carBounds.y) {
                car.y += carBounds.height;
            } else if (car.y > carBounds.y + carBounds.height) {
                car.y -= carBounds.height;
            }
        }
    });

}
$(() => {
    $(document).ready(function() {
        PikachuInit()

        Cars(25)
    })
});
