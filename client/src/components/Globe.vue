<template>
  <div id="chartContainer" ref="chartdiv"></div>
</template>

<script>
import * as am4core from "@amcharts/amcharts4/core";
import * as am4maps from "@amcharts/amcharts4/maps";
import worldGeoData from "@amcharts/amcharts4-geodata/worldLow";
import usaGeoData from "@amcharts/amcharts4-geodata/usaLow";
import indiaGeoData from "@amcharts/amcharts4-geodata/indiaLow";
import { getCountries } from "../api/map-api";

const visitedCountries = [
  "BE",
  "FR",
  "DE",
  "NL",
  "IT",
  "GB",
  "CA",
  "SG",
  "AT",
  "ES",
  "PL",
  "CH",
  "EG",
  "GR",
  "DO",
  "MA",
  "TR",
  "PT",
  "LU",
  "MC"
];
const visitedStatesUSA = [
  "US-CA",
  "US-TX",
  "US-FL",
  "US-AK",
  "US-WA",
  "US-NV",
  "US-MA",
  "US-OH",
  "US-WI",
  "US-CO",
  "US-IN",
  "US-IL",
  "US-NJ",
  "US-MN",
  "US-VA",
  "US-NY",
  "US-MO"
];
const visitedStatesIndia = [
  "IN-GJ",
  "IN-MH",
  "IN-GA",
  "IN-DL",
  "IN-UP",
  "IN-PB",
  "IN-RJ",
  "IN-KA"
];
const latitudeMap = new Map([
  ["IN", -80],
  ["US", 100],
  ["BE", 0],
  ["FR", 0],
  ["DE", -10],
  ["NL", 0],
  ["IT", -10],
  ["GB", 0],
  ["CA", 100],
  ["SG", -100],
  ["AT", -10],
  ["ES", 0],
  ["PL", -10],
  ["CH", 0],
  ["EG", -30],
  ["GR", -10],
  ["DO", 80],
  ["MA", 0],
  ["TR", -30],
  ["PT", 0],
  ["LU", 0],
  ["MC", 0]
]);

export default {
  name: "Globe",
  data: function() {
    return {
      chart: undefined,
      countrySeries: undefined,
      animation: undefined,
      visitedCountries: visitedCountries,
      visitedStatesUSA: visitedStatesUSA,
      visitedStatesIndia: visitedStatesIndia,
      latitudeMap: latitudeMap
    };
  },
  methods: {
    initGlobeState() {
      // map_api
      // get countries
      let countries = getCountries();
      console.log("countries");
      console.log(countries);
    },
    setupChart() {
      this.chart = am4core.create(this.$refs.chartdiv, am4maps.MapChart);
      this.chart.width = am4core.percent(100);
      this.chart.height = am4core.percent(100);
      this.chart.panBehavior = "rotateLongLat";

      // limits vertical rotation
      this.chart.adapter.add("deltaLatitude", function(delatLatitude) {
        return am4core.math.fitToRange(delatLatitude, -90, 90);
      });

      this.chart.geodata = worldGeoData;
      this.chart.projection = new am4maps.projections.Orthographic();

      this.chart.homeZoomLevel = 0;
      var home = this.chart.chartContainer.createChild(am4core.Button);
      home.label.text = "Home";
      home.align = "right";
      home.events.on("hit", () => {
        this.chart.goHome();
        this.animate();
      });
    },
    plotBgAndLines() {
      var latLongSeries = this.chart.series.push(new am4maps.GraticuleSeries());
      latLongSeries.mapLines.template.line.stroke = am4core.color("#67b7dc");
      latLongSeries.mapLines.template.line.strokeOpacity = 0.2;
      latLongSeries.fitExtent = false;

      // Background -> water
      this.chart.backgroundSeries.mapPolygons.template.polygon.fill = am4core.color(
        "#aadaff"
      );
      this.chart.backgroundSeries.mapPolygons.template.polygon.fillOpacity = 1;
    },
    plotCountries() {
      var disabledCountrySeries = this.chart.series.push(
        new am4maps.MapPolygonSeries()
      );
      disabledCountrySeries.north = 90;
      disabledCountrySeries.useGeodata = true; // load country data (names)
      disabledCountrySeries.exclude = visitedCountries;
      var notVisitedTemplate = disabledCountrySeries.mapPolygons.template;
      // notVisitedTemplate.tooltipText = "{name}";
      notVisitedTemplate.fill = am4core.color("#404F24"); // old color: #74B266
      notVisitedTemplate.opacity = 0.2;

      this.countrySeries = this.chart.series.push(
        new am4maps.MapPolygonSeries()
      );
      this.countrySeries.north = 90;
      this.countrySeries.useGeodata = true;
      this.countrySeries.include = visitedCountries;
      var visitedTemplate = this.countrySeries.mapPolygons.template;
      visitedTemplate.tooltipText = "{name}";
      visitedTemplate.fill = am4core.color("#74B266"); // old color: #74B266
      visitedTemplate.stroke = am4core.color("#000000"); // outline color
      visitedTemplate.strokeWidth = 0.5;
      visitedTemplate.strokeOpacity = 0.8;
      visitedTemplate.events.on("hit", ev => {
        if (this.animation) {
          this.animation.stop();
        }
        // zoom to an object
        ev.target.series.chart.zoomToMapObject(ev.target);

        // get object info
        console.log(ev.target);
      });
      var hoverState = visitedTemplate.states.create("hover");
      hoverState.properties.fill = am4core.color("#367B25");
    },
    plotStates(geoData, visitedStates) {
      let stateSeries = this.chart.series.push(new am4maps.MapPolygonSeries());
      stateSeries.geodata = geoData;
      stateSeries.include = visitedStates;
      let statesTemplate = stateSeries.mapPolygons.template;
      statesTemplate.tooltipText = "{name}";
      // statesTemplate.fill = this.chart.colors.getIndex(1);
      // statesTemplate.nonScalingStroke = true;
      statesTemplate.fill = am4core.color("#74B266"); // old color: #74B266
      statesTemplate.stroke = am4core.color("#000000"); // outline color
      statesTemplate.strokeWidth = 0.5;
      statesTemplate.strokeOpacity = 0.8;
      statesTemplate.events.on("hit", ev => {
        if (this.animation) {
          this.animation.stop();
        }
        // zoom to an object
        ev.target.series.chart.zoomToMapObject(ev.target);

        // get object info
        console.log(ev.target);
      });
      let hoverState = statesTemplate.states.create("hover");
      hoverState.properties.fill = am4core.color("#367B25");
    },
    focusOnCountry(polygonSeries, countryCode) {
      // TODO: error or return on else? or leave behaviour as is?
      if (this.latitudeMap.has(countryCode)) {
        this.chart.deltaLongitude = this.latitudeMap.get(countryCode);
      }
      this.chart.events.on("ready", () => {
        // NOTE: this will only work when the polygon is included in the series (USA and INDIA -> removed from global series)
        // this.chart.events.on("ready", function() {
        var countryPolygon = polygonSeries.getPolygonById(countryCode);
        this.chart.zoomToMapObject(countryPolygon);

        // Set active state
        // setTimeout(function() {
        //   countryPolygon.isActive = true;
        // }, 1000);
      });
    },
    rotateTo(long, lat) {
      if (this.animation) {
        this.animation.stop();
      }
      this.animation = this.chart.animate(
        [
          {
            property: "deltaLongitude",
            to: long
          },
          {
            property: "deltaLatitude",
            to: lat
          }
        ],
        2000
      );
    },
    animate() {
      // this.chart.goHome(); // to reset zoom
      this.animation = this.chart.animate(
        { property: "deltaLongitude", to: 100000 },
        20000000
      );

      this.chart.seriesContainer.events.on("down", () => {
        if (this.animation) {
          this.animation.stop();
        }
      });
    }
  },
  mounted: function() {
    this.initGlobeState();
    // Create map instance
    this.setupChart();
    this.plotBgAndLines();
    this.plotCountries();
    this.plotStates(usaGeoData, this.visitedStatesUSA);
    this.plotStates(indiaGeoData, this.visitedStatesIndia);
    this.focusOnCountry(this.countrySeries, "SG"); // TODO: move out of here (testing)
    // this.animate();
  }
};
</script>

<style scoped>
#chartContainer {
  width: 75%;
  height: 75%;
  position: absolute;
  margin: auto;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  overflow: hidden;
  background: #000;
}
</style>
