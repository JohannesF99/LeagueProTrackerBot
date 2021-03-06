package model

type ChallengesDTO struct {
	AssistStreakCount                        int     `json:"12AssistStreakCount"`
	AbilityUses                              int     `json:"abilityUses"`
	AcesBefore15Minutes                      int     `json:"acesBefore15Minutes"`
	AlliedJungleMonsterKills                 float64 `json:"alliedJungleMonsterKills"`
	BaronTakedowns                           int     `json:"baronTakedowns"`
	BlastConeOppositeOpponentCount           int     `json:"blastConeOppositeOpponentCount"`
	BountyGold                               int     `json:"bountyGold"`
	BuffsStolen                              int     `json:"buffsStolen"`
	CompleteSupportQuestInTime               int     `json:"completeSupportQuestInTime"`
	ControlWardsPlaced                       int     `json:"controlWardsPlaced"`
	DamagePerMinute                          float64 `json:"damagePerMinute"`
	DamageTakenOnTeamPercentage              float64 `json:"damageTakenOnTeamPercentage"`
	DancedWithRiftHerald                     int     `json:"dancedWithRiftHerald"`
	DeathsByEnemyChamps                      int     `json:"deathsByEnemyChamps"`
	DodgeSkillShotsSmallWindow               int     `json:"dodgeSkillShotsSmallWindow"`
	DoubleAces                               int     `json:"doubleAces"`
	DragonTakedowns                          int     `json:"dragonTakedowns"`
	EarlyLaningPhaseGoldExpAdvantage         float64 `json:"earlyLaningPhaseGoldExpAdvantage"`
	EffectiveHealAndShielding                float64 `json:"effectiveHealAndShielding"`
	ElderDragonKillsWithOpposingSoul         int     `json:"elderDragonKillsWithOpposingSoul"`
	ElderDragonMultikills                    int     `json:"elderDragonMultikills"`
	EnemyChampionImmobilizations             int     `json:"enemyChampionImmobilizations"`
	EnemyJungleMonsterKills                  float64 `json:"enemyJungleMonsterKills"`
	EpicMonsterKillsNearEnemyJungler         int     `json:"epicMonsterKillsNearEnemyJungler"`
	EpicMonsterKillsWithin30SecondsOfSpawn   int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
	EpicMonsterSteals                        int     `json:"epicMonsterSteals"`
	EpicMonsterStolenWithoutSmite            int     `json:"epicMonsterStolenWithoutSmite"`
	FlawlessAces                             int     `json:"flawlessAces"`
	FullTeamTakedown                         int     `json:"fullTeamTakedown"`
	GameLength                               float64 `json:"gameLength"`
	GetTakedownsInAllLanesEarlyJungleAsLaner int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner"`
	GoldPerMinute                            float64 `json:"goldPerMinute"`
	HadAfkTeammate                           int     `json:"hadAfkTeammate"`
	HadOpenNexus                             int     `json:"hadOpenNexus"`
	InitialBuffCount                         int     `json:"initialBuffCount"`
	InitialCrabCount                         int     `json:"initialCrabCount"`
	JungleCsBefore10Minutes                  float64 `json:"jungleCsBefore10Minutes"`
	JunglerKillsEarlyJungle                  int     `json:"junglerKillsEarlyJungle"`
	JunglerTakedownsNearDamagedEpicMonster   int     `json:"junglerTakedownsNearDamagedEpicMonster"`
	KTurretsDestroyedBeforePlatesFall        int     `json:"kTurretsDestroyedBeforePlatesFall"`
	Kda                                      float64 `json:"kda"`
	KillParticipation                        float64 `json:"killParticipation"`
	KillsNearEnemyTurret                     int     `json:"killsNearEnemyTurret"`
	KillsOnLanersEarlyJungleAsJungler        int     `json:"killsOnLanersEarlyJungleAsJungler"`
	KillsOnOtherLanesEarlyJungleAsLaner      int     `json:"killsOnOtherLanesEarlyJungleAsLaner"`
	KillsOnRecentlyHealedByAramPack          int     `json:"killsOnRecentlyHealedByAramPack"`
	KillsUnderOwnTurret                      int     `json:"killsUnderOwnTurret"`
	KillsWithHelpFromEpicMonster             int     `json:"killsWithHelpFromEpicMonster"`
	LandSkillShotsEarlyGame                  int     `json:"landSkillShotsEarlyGame"`
	LaneMinionsFirst10Minutes                int     `json:"laneMinionsFirst10Minutes"`
	LaningPhaseGoldExpAdvantage              float64 `json:"laningPhaseGoldExpAdvantage"`
	LegendaryCount                           int     `json:"legendaryCount"`
	LostAnInhibitor                          int     `json:"lostAnInhibitor"`
	MaxCsAdvantageOnLaneOpponent             float64 `json:"maxCsAdvantageOnLaneOpponent"`
	MaxKillDeficit                           int     `json:"maxKillDeficit"`
	MaxLevelLeadLaneOpponent                 int     `json:"maxLevelLeadLaneOpponent"`
	MoreEnemyJungleThanOpponent              float64 `json:"moreEnemyJungleThanOpponent"`
	MultiKillOneSpell                        int     `json:"multiKillOneSpell"`
	MultiTurretRiftHeraldCount               int     `json:"multiTurretRiftHeraldCount"`
	Multikills                               int     `json:"multikills"`
	MultikillsAfterAggressiveFlash           int     `json:"multikillsAfterAggressiveFlash"`
	MythicItemUsed                           int     `json:"mythicItemUsed"`
	OuterTurretExecutesBefore10Minutes       int     `json:"outerTurretExecutesBefore10Minutes"`
	OutnumberedKills                         int     `json:"outnumberedKills"`
	OutnumberedNexusKill                     int     `json:"outnumberedNexusKill"`
	PerfectDragonSoulsTaken                  int     `json:"perfectDragonSoulsTaken"`
	PerfectGame                              int     `json:"perfectGame"`
	PoroExplosions                           int     `json:"poroExplosions"`
	QuickCleanse                             int     `json:"quickCleanse"`
	QuickFirstTurret                         int     `json:"quickFirstTurret"`
	QuickSoloKills                           int     `json:"quickSoloKills"`
	RiftHeraldTakedowns                      int     `json:"riftHeraldTakedowns"`
	ScuttleCrabKills                         int     `json:"scuttleCrabKills"`
	SkillshotsDodged                         int     `json:"skillshotsDodged"`
	SkillshotsHit                            int     `json:"skillshotsHit"`
	SnowballsHit                             int     `json:"snowballsHit"`
	SoloBaronKills                           int     `json:"soloBaronKills"`
	SoloKills                                int     `json:"soloKills"`
	SoloTurretsLategame                      int     `json:"soloTurretsLategame"`
	StealthWardsPlaced                       int     `json:"stealthWardsPlaced"`
	SurvivedSingleDigitHpCount               int     `json:"survivedSingleDigitHpCount"`
	TakedownOnFirstTurret                    int     `json:"takedownOnFirstTurret"`
	Takedowns                                int     `json:"takedowns"`
	TakedownsAfterGainingLevelAdvantage      int     `json:"takedownsAfterGainingLevelAdvantage"`
	TakedownsBeforeJungleMinionSpawn         int     `json:"takedownsBeforeJungleMinionSpawn"`
	TakedownsFirst25Minutes                  int     `json:"takedownsFirst25Minutes"`
	TakedownsInAlcove                        int     `json:"takedownsInAlcove"`
	TakedownsInEnemyFountain                 int     `json:"takedownsInEnemyFountain"`
	TeamBaronKills                           int     `json:"teamBaronKills"`
	TeamDamagePercentage                     float64 `json:"teamDamagePercentage"`
	TeamElderDragonKills                     int     `json:"teamElderDragonKills"`
	TeamRiftHeraldKills                      int     `json:"teamRiftHeraldKills"`
	TeleportTakedowns                        int     `json:"teleportTakedowns"`
	ThreeWardsOneSweeperCount                int     `json:"threeWardsOneSweeperCount"`
	TurretPlatesTaken                        int     `json:"turretPlatesTaken"`
	TurretTakedowns                          int     `json:"turretTakedowns"`
	TurretsTakenWithRiftHerald               int     `json:"turretsTakenWithRiftHerald"`
	TwentyMinionsIn3SecondsCount             int     `json:"twentyMinionsIn3SecondsCount"`
	UnseenRecalls                            int     `json:"unseenRecalls"`
	VisionScoreAdvantageLaneOpponent         float64 `json:"visionScoreAdvantageLaneOpponent"`
	VisionScorePerMinute                     float64 `json:"visionScorePerMinute"`
	WardTakedowns                            int     `json:"wardTakedowns"`
	WardTakedownsBefore20M                   int     `json:"wardTakedownsBefore20M"`
	WardsGuarded                             int     `json:"wardsGuarded"`
}
