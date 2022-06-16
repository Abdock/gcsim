package diona

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

func (c *char) c2() {
	c.AddAttackMod("diona-c2",
		-1,
		func(atk *combat.AttackEvent, t combat.Target) ([]float64, bool) {
			val := make([]float64, attributes.EndStatType)
			val[attributes.DmgP] = .15
			return val, atk.Info.AttackTag == combat.AttackTagElementalArt
		})

}
func (c *char) c6() {
	//c6 should last for the duration of the burst

	//lasts 12.5 second, ticks every 0.5s; adds mod to active char for 2s
	for i := 30; i < 750; i += 30 {
		c.Core.Tasks.Add(func() {
			//add 200EM to active char
			char := c.Core.Player.ActiveChar()
			if char.HPCurrent/char.MaxHP() > 0.5 {
				char.AddStatMod(
					"diona-c6",
					120, //lasts 2 seconds
					attributes.NoStat,
					func() ([]float64, bool) {
						return c.c6buff, true
					},
				)
			} else {
				//add healing bonus if hp <= 0.5
				//bonus only lasts for 120 frames
				char.AddHealBonusMod(
					"diona-c6-healbonus",
					120,
					func() (float64, bool) {
						c.Core.Log.NewEvent("diona c6 incomming heal bonus activated", glog.LogCharacterEvent, c.Index)
						return 0.3, false
					},
				)
				c.Tags["c6bonus-"+char.Base.Key.String()] = c.Core.F + 120
			}
		}, i)
	}
}
