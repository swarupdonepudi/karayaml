import React from "react";
import Hero from "../components/home/Hero";
import FeatureCards from "../components/home/FeatureCards";
import HowItWorks from "../components/home/HowItWorks";
import Quickstart from "../components/home/Quickstart";
import ExampleGallery from "../components/home/ExampleGallery";
import CLIReference from "../components/home/CLIReference";
import KeyUsageInfo from "../components/home/KeyUsageInfo";
import Comparison from "../components/home/Comparison";
import FAQ from "../components/home/FAQ";

export default function Home() {
  return (
    <div className="overflow-hidden">
      <Hero />
      <FeatureCards />
      <HowItWorks />
      <Quickstart />
      <ExampleGallery />
      <CLIReference />
      <KeyUsageInfo />
      <Comparison />
      <FAQ />
    </div>
  );
}