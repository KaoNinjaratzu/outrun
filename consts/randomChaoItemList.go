package consts

import "github.com/KaoNinjaratzu/outrun/enums"

type PrizeInfo struct {
	AppearanceChance float64 // % chance for it to be chosen to be in wheel by the server
	Type             int64   // 0 for Chao, 1 for Character
}

// A 'load' as depicted below is the chance for the server to pick
// the associated item, where chosen is if randFloat(0, 100) < load.
// IMPORTANT: This load is exclusive to the rarity of the Chao that
// is being chosen by the server.

var RandomChaoWheelCharacterPrizes = map[string]float64{
	// characterID: load
	// Hopefully this should sum up to 100 just for
	// simplicity, but it shouldn't be a requirement.
	enums.CTStrSonic:        1.0,
	enums.CTStrTails:        1.0,
	enums.CTStrKnuckles:     1.0,
	enums.CTStrAmy:          1.0,
	enums.CTStrBig:          0.7,
	//enums.CTStrBlaze:        0.0, // Revival Event
	enums.CTStrCharmy:       1.0,
	enums.CTStrCream:        0.7,
	enums.CTStrEspio:        1.0,
	//enums.CTStrMephiles:     0.0, // Revival Event
	enums.CTStrOmega:        0.5,
	//enums.CTStrPSISilver:    0.0, // Revival Event
	enums.CTStrRouge:        0.5,
	enums.CTStrShadow:       0.5,
	enums.CTStrSilver:       0.7,
	//enums.CTStrSticks:       0.0, // Revival Event
	//enums.CTStrTikal:        0.0, // Event (Sonic Adventure)
	enums.CTStrVector:       1.0,
	//enums.CTStrWerehog:      0.0, // Revival Event
	//enums.CTStrClassicSonic: 0.0, // Event (Birthday)
	//enums.CTStrMetalSonic:   0.0, // Revival Event
}

var RandomChaoWheelChaoPrizes = map[string]float64{
	// TODO: Balance these
	//enums.ChaoIDStrHeroChao:             0.0, // Event (Animal Rescue event)
	//enums.ChaoIDStrGoldChao:             0.0, // Event (Animal Rescue event)
	//enums.ChaoIDStrDarkChao:             0.0, // Event (Animal Rescue event)
	//enums.ChaoIDStrJewelChao:            0.0, // Event (Animal Rescue event)
	enums.ChaoIDStrNormalChao:           3.0,
	enums.ChaoIDStrOmochao:              3.0,
	//enums.ChaoIDStrRCMonkey:             0.0, // Event (Animal Rescue event)
	enums.ChaoIDStrRCSpring:             2.0,
	enums.ChaoIDStrRCElectromagnet:      2.0,
	enums.ChaoIDStrBabyCyanWisp:         3.0,
	enums.ChaoIDStrBabyIndigoWisp:       3.0,
	enums.ChaoIDStrBabyYellowWisp:       3.0,
	enums.ChaoIDStrRCPinwheel:           2.0,
	enums.ChaoIDStrRCPiggyBank:          1.0,
	enums.ChaoIDStrRCBalloon:            2.0,
	//enums.ChaoIDStrEasterChao:           0.0, // Event (Easter)
	//enums.ChaoIDStrPurplePapurisu:       0.0, // Event (Puyo Puyo)
	//enums.ChaoIDStrMagLv1:               0.0, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrEggChao:              1.0,
	enums.ChaoIDStrPumpkinChao:          1.0,
	enums.ChaoIDStrSkullChao:            1.0,
	enums.ChaoIDStrYacker:               1.0,
	enums.ChaoIDStrRCGoldenPiggyBank:    1.0,
	enums.ChaoIDStrWizardChao:           1.0,
	enums.ChaoIDStrRCTurtle:             1.0,
	enums.ChaoIDStrRCUFO:                1.0,
	enums.ChaoIDStrRCBomber:             1.0,
	//enums.ChaoIDStrEasterBunny:          0.0, // Event (Easter)
	//enums.ChaoIDStrMagicLamp:            0.0, // Event (Desert)
	//enums.ChaoIDStrStarShapedMissile:    0.0, // Event (Raid Boss Roulette)
	//enums.ChaoIDStrSuketoudara:          0.0, // Event (Puyo Puyo)
	//enums.ChaoIDStrRappy:                0.0, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrBlowfishTransporter:  0.0, // Event (Beach)
	//enums.ChaoIDStrGenesis:              0.0, // Event (Birthday)
	//enums.ChaoIDStrCartridge:            0.0, // Event (Birthday)
	enums.ChaoIDStrRCFighter:            1.0,
	enums.ChaoIDStrRCHovercraft:         1.0,
	enums.ChaoIDStrRCHelicopter:         1.0,
	enums.ChaoIDStrGreenCrystalMonsterS: 1.0,
	enums.ChaoIDStrGreenCrystalMonsterL: 1.0,
	enums.ChaoIDStrRCAirship:            1.0,
	//enums.ChaoIDStrDesertChao:           0.0, // Event (Desert)
	//enums.ChaoIDStrRCSatellite:          0.0, // Event (Raid Boss Roulette?)
	//enums.ChaoIDStrMarineChao:           0.0, // Event (Beach)
	//enums.ChaoIDStrNightopian:           0.0, // Event (NiGHTS)
	//enums.ChaoIDStrOrca:                 0.0, // Event (Sonic Adventure)
	//enums.ChaoIDStrSonicOmochao:         0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrTailsOmochao:         0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrKnucklesOmochao:      0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrBoo:                  0.0, // Event (Halloween)
	//enums.ChaoIDStrHalloweenChao:        0.0, // Event (Halloween)
	//enums.ChaoIDStrHeavyBomb:            0.0, // Event (Fantasy Zone)
	enums.ChaoIDStrBlockBomb:            1.0,
	enums.ChaoIDStrHunkofMeat:           1.0,
	//enums.ChaoIDStrYeti:                 0.0, // Event (Christmas)
	//enums.ChaoIDStrSnowChao:             0.0, // Event (Christmas)
	//enums.ChaoIDStrIdeya:                0.0, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrChristmasNightopian:  0.0, // Event (Christmas NiGHTS)
	enums.ChaoIDStrOrbot:                1.0,
	enums.ChaoIDStrCubot:                1.0,
	enums.ChaoIDStrLightChaos:           1.5,
	enums.ChaoIDStrHeroChaos:            1.5,
	enums.ChaoIDStrDarkChaos:            1.5,
	enums.ChaoIDStrChip:                 1.5,
	//enums.ChaoIDStrShahra:               0.2, // Runners League
	enums.ChaoIDStrCaliburn:             1.5,
	enums.ChaoIDStrKingArthursGhost:     1.5,
	enums.ChaoIDStrRCTornado:            1.0,
	enums.ChaoIDStrRCBattleCruiser:      1.0,
	//enums.ChaoIDStrMerlina:              0.0, // Event (Windy Hill in Spring)
	//enums.ChaoIDStrErazorDjinn:          0.0, // Event (Desert Ruins)
	//enums.ChaoIDStrRCMoonMech:           0.0, // Raid Boss Roulette(?)
	//enums.ChaoIDStrCarbuncle:            0.0, // Event (Puyo Puyo)
	//enums.ChaoIDStrKuna:                 0.0, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrChaos:                0.0, // Event (Sonic Adventure)
	//enums.ChaoIDStrDeathEgg:             0.0, // Event (Birthday)
	enums.ChaoIDStrRedCrystalMonsterS:   1.0,
	enums.ChaoIDStrRedCrystalMonsterL:   1.0,
	enums.ChaoIDStrGoldenGoose:          1.0,
	//enums.ChaoIDStrMotherWisp:           0.0, // Event (Tropical Coast)
	enums.ChaoIDStrRCPirateSpaceship:    1.0,
	enums.ChaoIDStrGoldenAngel:          1.0,
	//enums.ChaoIDStrNiGHTS:               0.0, // Event (NiGHTS)
	//enums.ChaoIDStrReala:                0.0, // Event (NiGHTS)
	//enums.ChaoIDStrRCTornado2:           0.0, // Event (Sonic Adventure
	enums.ChaoIDStrChaoWalker:           1.0,
	//enums.ChaoIDStrDarkQueen:            0.2, // Runners Leauge
	//enums.ChaoIDStrKingBoomBoo:          0.0, // Event (Halloween)
	//enums.ChaoIDStrOPapa:                0.0, // Event (Fantasy Zone)
	//enums.ChaoIDStrOpaOpa:               0.0, // Event (Fantasy Zone)
	enums.ChaoIDStrRCBlockFace:          1.0,
	//enums.ChaoIDStrChristmasYeti:        0.0, // Event (Christmas)
	//enums.ChaoIDStrChristmasNiGHTS:      0.0, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrDFekt:                0.0, // DO NOT ENABLE
	enums.ChaoIDStrDarkChaoWalker:       1.0,
}
