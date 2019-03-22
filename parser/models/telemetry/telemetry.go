package telemetry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
)

// TelemetryEventType represents the type of a telemetry event
type EventType int

func findIndex(key string, possibleKeys []string) int {
	for idx, possibleKey := range possibleKeys {
		if key == possibleKey {
			return idx
		}
	}
	return -1
}

// Telemetry event types
const (
	PlayerLogin EventType = iota
	PlayerLogout
	PlayerCreate
	PlayerPosition
	PlayerAttack
	PlayerTakeDamage
	PlayerKill
	ItemPickup
	ItemDrop
	ItemEquip
	ItemUnequip
	ItemAttach
	ItemDetach
	ItemUse
	VehicleRide
	VehicleLeave
	VehicleDestroy
	MatchStart
	MatchEnd
	MatchDefinition
	GameStatePeriodic
	CarePackageSpawn
	CarePackageLand
	VaultStart
	SwimStart
	ParachuteLanding
	PlayerMakeGroggy
	WeaponFireCount
	ArmorDestroy
	ItemPickupFromLootBox
	Heal
	ObjectDestroy
	PlayerRevive
	SwimEnd
	WheelDestroy
	RedZoneEnded
)

// KnownEventTypes represents supported types
var KnownEventTypes = []string{
	"LogPlayerLogin",
	"LogPlayerLogout",
	"LogPlayerCreate",
	"LogPlayerPosition",
	"LogPlayerAttack",
	"LogPlayerTakeDamage",
	"LogPlayerKill",
	"LogItemPickup",
	"LogItemDrop",
	"LogItemEquip",
	"LogItemUnequip",
	"LogItemAttach",
	"LogItemDetach",
	"LogItemUse",
	"LogVehicleRide",
	"LogVehicleLeave",
	"LogVehicleDestroy",
	"LogMatchStart",
	"LogMatchEnd",
	"LogMatchDefinition",
	"LogGameStatePeriodic",
	"LogCarePackageSpawn",
	"LogCarePackageLand",
	"LogVaultStart",
	"LogSwimStart",
	"LogParachuteLanding",
	"LogPlayerMakeGroggy",
	"LogWeaponFireCount",
	"LogArmorDestroy",
	"LogItemPickupFromLootBox",
	"LogHeal",
	"LogObjectDestroy",
	"LogPlayerRevive",
	"LogSwimEnd",
	"LogWheelDestroy",
	"LogRedZoneEnded",
}

// UnmarshalJSON verifies the type of a telemetry event is known
func (t *EventType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryEventType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], KnownEventTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryEventType: Unknown type %s", key)
	}

	*t = EventType(idx)
	return nil
}

// TelemetryAttackType represents the type of an attack
type AttackType int

// Telemetry attack types
const (
	AttackTypeRedZone AttackType = iota
	AttackTypeWeapon
)

var knownAttackTypes = []string{
	"RedZone",
	"Weapon",
}

// UnmarshalJSON verifies the type of an attack is known
func (t *AttackType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryAttackType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownAttackTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryAttackType: Unknown type %s", key)
	}

	*t = AttackType(idx)
	return nil
}

// TelemetrySubCategory represents the category of an item
type SubCategory int

// Telemetry sub categories
const (
	SubCategoryBackpack SubCategory = iota
	SubCategoryBoost
	SubCategoryFuel
	SubCategoryHandgun
	SubCategoryHeadgear
	SubCategoryHeal
	SubCategoryMain
	SubCategoryMelee
	SubCategoryThrowable
	SubCategoryVest
	SubCategoryJacket
	SubCategoryNone
	SubCategoryEmpty
)

var knownSubCategories = []string{
	"Backpack",
	"Boost",
	"Fuel",
	"Handgun",
	"Headgear",
	"Heal",
	"Main",
	"Melee",
	"Throwable",
	"Vest",
	"Jacket",
	"None",
	"",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *SubCategory) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetrySubCategory: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownSubCategories)
	if idx == -1 {
		return fmt.Errorf("TelemetrySubCategory: Unknown type %s", key)
	}

	*t = SubCategory(idx)
	return nil
}

// TelemetryDamageType represents the different types of damage
type DamageType int

// Telemetry damage types
const (
	DamageBlueZone DamageType = iota
	DamageDrown
	DamageExplosionGrenade
	DamageExplosionRedZone
	DamageExplosionVehicle
	DamageGroggy
	DamageGun
	DamageInstantFall
	DamageMelee
	DamageMolotov
	DamageVehicleCrashHit
	DamageVehicleHit
	DamageEmpty
)

var knownDamageTypes = []string{
	"Damage_BlueZone",
	"Damage_Drown",
	"Damage_Explosion_Grenade",
	"Damage_Explosion_RedZone",
	"Damage_Explosion_Vehicle",
	"Damage_Groggy",
	"Damage_Gun",
	"Damage_Instant_Fall",
	"Damage_Melee",
	"Damage_Molotov",
	"Damage_VehicleCrashHit",
	"Damage_VehicleHit",
	"Damage_Punch",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *DamageType) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryDamageType: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownDamageTypes)
	if idx == -1 {
		return fmt.Errorf("TelemetryDamageType: Unknown type %s", key)
	}

	*t = DamageType(idx)
	return nil
}

// TelemetryDamageReason represents the reason of the damage
type DamageReason int

// Telemetry damage reasons
const (
	DamageReasonArmShot DamageReason = iota
	DamageReasonHeadShot
	DamageReasonLegShot
	DamageReasonPelvisShot
	DamageReasonTorsoShot
	DamageReasonNonSpecific
	DamageReasonNone
)

var knownDamageReasons = []string{
	"ArmShot",
	"HeadShot",
	"LegShot",
	"PelvisShot",
	"TorsoShot",
	"NonSpecific",
	"None",
	"",
}

// UnmarshalJSON verifies the type of a subcategory
func (t *DamageReason) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("TelemetryDamageReason: UnmarshalJSON on nil pointer")
	}

	key := string(data)
	idx := findIndex(key[1:len(key)-1], knownDamageReasons)
	if idx == -1 {
		return fmt.Errorf("TelemetryDamageReason: Unknown type %s", key)
	}

	*t = DamageReason(idx)
	return nil
}

// TelemetryEvent represents any event from a telemetry file
type Event struct {
	// Common fields
	//Version   int                `json:"_V"`
	//Timestamp time.Time          `json:"_D"`
	Type EventType `json:"_T"`
	//U         bool               `json:"_U"`

	// --- Player
	// Events: LogPlayerLogin, LogPlayerLogout, LogPlayerCreate, LogPlayerPosition, LogPlayerAttack, LogPlayerTakeDamage, LogPlayerKill
	//Result             bool                  `json:"result"`
	//ErrorMessage       string                `json:"errorMessage"`
	//AccountID          string                `json:"accountId"`
	//Character          *TelemetryCharacter   `json:"character"`
	//ElapsedTime        float64               `json:"elapsedTime"`
	//NumAlivePlayers    int                   `json:"numAlivePlayers"`
	AttackID int `json:"attackId"`
	//Attacker           *TelemetryCharacter   `json:"attacker"`
	//AttackType         TelemetryAttackType   `json:"attackType"`
	//Weapon             *TelemetryItem        `json:"weapon"`
	//Vehicle            *TelemetryVehicle     `json:"vehicle"`
	Victim *Character `json:"victim"`
	//DamageTypeCategory TelemetryDamageType   `json:"damageTypeCategory"`
	//DamageReason       TelemetryDamageReason `json:"damageReason"`
	//Damage             float64               `json:"damage"`
	DamageCauserName string     `json:"damageCauserName"`
	Killer           *Character `json:"killer"`
	Distance         float64    `json:"distance"`

	// --- Vehicle
	// Events: LogVehicleRide, LogVehicleLeave, VehicleDestroy
	// Character already defined
	// Vehicle already defined

	// --- Item
	// Events: LogItemPickup, LogItemEquip, LogItemUnequip, LogItemAttach, LogItemDrop, LogItemDetach, LogItemUse
	//Item       *TelemetryItem `json:"item"`
	//ParentItem *TelemetryItem `json:"parentItem"`
	//ChildItem  *TelemetryItem `json:"childItem"`

	// --- Match
	// Events: LogMatchStart, LogMatchEnd, LogMatchDefinition
	//Characters  []*TelemetryCharacter
	MatchID string `json:"matchId"`
	//PingQuality string `json:"pingQuality"`

	// --- Care package
	// Events: LogCarePackageSpawn, LogCarePackageLand
	//ItemPackage *TelemetryItemPackage `json:"itemPackage"`

	// --- Game
	// Events: LogGameStatePeriodic
	GameState *GameState
}

// TelemetryItemPackage represents an item package
type ItemPackage struct {
	ItemPackageID string    `json:"itemPackageId"`
	Location      *Location `json:"location"`
	Items         []*Item   `json:"items"`
}

// TelemetryGameState represents the state of a game
type GameState struct {
	ElapsedTime              int       `json:"elapsedTime"`
	NumAliveTeams            int       `json:"numAliveTeams"`
	NumJoinPlayers           int       `json:"numJoinPlayers"`
	NumStartPlayers          int       `json:"numStartPlayers"`
	NumAlivePlayers          int       `json:"numAlivePlayers"`
	SafetyZonePosition       *Location `json:"safetyZonePosition"`
	SafetyZoneRadius         float64   `json:"safetyZoneRadius"`
	PoisonGasWarningPosition *Location `json:"poisonGasWarningPosition"`
	PoisonGasWarningRadius   float64   `json:"poisonGasWarningRadius"`
	RedZonePosition          *Location `json:"redZonePosition"`
	RedZoneRadius            float64   `json:"redZoneRadius"`
}

// TelemetryVehicle represents a vehicle
type Vehicle struct {
	VehicleType   string  `json:"vehicleType"`
	VehicleID     string  `json:"vehicleId"`
	HealthPercent float64 `json:"healthPercent"`
	FuelPercent   float64 `json:"feulPercent"`
}

// TelemetryItem represents an item
type Item struct {
	ItemID        string      `json:"itemId"`
	StackCount    int         `json:"stackCount"`
	Category      string      `json:"category"`
	SubCategory   SubCategory `json:"subCategory"`
	AttachedItems []string    `json:"attachedItems"`
}

// TelemetryCharacter represents a character
type Character struct {
	Name      string    `json:"name"`
	TeamID    int       `json:"teamId"`
	Health    float64   `json:"health"`
	Location  *Location `json:"location"`
	Ranking   int       `json:"ranking"`
	AccountID string    `json:"accountId"`
}

// TelemetryLocation represents a location
type Location struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
	Z float64 `json:"Z"`
}

// Player represents a player
type Player struct {
	Name      string
	AccountID string
	TeamID    int
	Events    []*Event
	Locations []*Location
	Ranking   int
}

func newPlayer(name, accountID string) *Player {
	return &Player{
		Name:      name,
		AccountID: accountID,
		TeamID:    -1,
		Events:    make([]*Event, 0),
		Locations: make([]*Location, 0),
		Ranking:   -1,
	}
}

// Telemetry represents the context of a telemetry file
type Telemetry struct {
	Events       []*Event
	Players      map[string]*Player
	MatchStarted bool
	PingQuality  string
	MatchID      string
}

func newTelemetry() *Telemetry {
	return &Telemetry{
		Events:       make([]*Event, 0),
		Players:      make(map[string]*Player),
		MatchStarted: false,
		PingQuality:  "",
		MatchID:      "",
	}
}

func (t *Telemetry) getPlayer(name, accountID string) *Player {
	if _, ok := t.Players[accountID]; !ok {
		t.Players[accountID] = newPlayer(name, accountID)
	}
	return t.Players[accountID]
}

func (t *Telemetry) addPlayerEvent(te *Event, character *Character, matchStarted bool) {
	if character == nil || character.Name == "" {
		return
	}
	fmt.Println(character.Name)
	player := t.getPlayer(character.Name, character.AccountID)
	if matchStarted {
		player.Events = append(player.Events, te)
		player.Locations = append(player.Locations, character.Location)
	}
}

func (t *Telemetry) processEvent(te *Event) {
	//logrus.WithFields(logrus.Fields{
	//	"type": KnownEventTypes[te.Type],
	//}).Debug("Processing event")

	// Look for common fields
	//if te.Character != nil {
	//	t.addPlayerEvent(te, te.Character, t.MatchStarted)
	//}

	// Look for a custom function to specialize the data
	functionName := "Process" + KnownEventTypes[te.Type]
	f := reflect.ValueOf(t).MethodByName(functionName)
	if f.IsValid() {
		f.Call([]reflect.Value{
			reflect.ValueOf(te),
		})
	}
}

// ProcessLogMatchDefinition deals with event of type MatchDefinition
//func (t *Telemetry) ProcessLogMatchDefinition(te *TelemetryEvent) {
//	t.PingQuality = te.PingQuality
//	t.MatchID = te.MatchID
//}

// ProcessLogMatchStart deals with event of type MatchStart
func (t *Telemetry) ProcessLogMatchStart(te *Event) {
	t.MatchStarted = true
}

// ProcessLogMatchEnd deals with event of type MatchEnd
//func (t *Telemetry) ProcessLogMatchEnd(te *TelemetryEvent) {
//	// Update player ranking
//	for _, c := range te.Characters {
//		player := t.getPlayer(c.Name, c.AccountID)
//		player.Ranking = c.Ranking
//	}
//}

// ProcessLogPlayerCreate deals with event of type PlayerCreate
//func (t *Telemetry) ProcessLogPlayerCreate(te *TelemetryEvent) {
//	player := t.getPlayer(te.Character.Name, te.Character.AccountID)
//	player.TeamID = te.Character.TeamID
//}

// ProcessLogPlayerAttack deals with event of type PlayerAttack
//func (t *Telemetry) ProcessLogPlayerAttack(te *TelemetryEvent) {
//	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
//}

// ProcessLogPlayerAttack deals with event of type PlayerAttack
func (t *Telemetry) ProcessLogPlayerKill(te *Event) {
	t.addPlayerEvent(te, te.Killer, t.MatchStarted)
	t.addPlayerEvent(te, te.Victim, t.MatchStarted)
}

// ProcessLogPlayerTakeDamage deals with event of type PlayerTakeDamage
//func (t *Telemetry) ProcessLogPlayerTakeDamage(te *TelemetryEvent) {
//	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
//	t.addPlayerEvent(te, te.Victim, t.MatchStarted)
//}

// ProcessLogVehicleDestroy deals with event of type VehicleDestroy
//func (t *Telemetry) ProcessLogVehicleDestroy(te *TelemetryEvent) {
//	t.addPlayerEvent(te, te.Attacker, t.MatchStarted)
//}

// ParseTelemetry parses a json response containing telemetry information
func ParseTelemetry(in io.Reader) (*Telemetry, error) {
	t := newTelemetry()
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	// Parse events
	if err := json.Unmarshal(data, &t.Events); err != nil {
		return nil, err
	}

	// Find players
	for _, e := range t.Events {
		t.processEvent(e)
	}

	return t, nil
}
