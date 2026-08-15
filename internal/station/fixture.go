package station

func fixtureCatalog() Catalog {
	return Catalog{
		StationName: "社区花园观察站",
		MonthLabel:  "八月",
		Plants: []Plant{
			{ID: "cherry-tomato", Name: "樱桃番茄", LatinName: "Solanum lycopersicum", Observation: "果穗进入连续成熟期，午后需要检查盆土湿度。", Image: Image{AssetID: "plant-cherry-tomato", Alt: "藤架上成熟的樱桃番茄"}},
			{ID: "purple-basil", Name: "紫罗勒", LatinName: "Ocimum basilicum", Observation: "新叶颜色稳定，摘心后侧芽数量增加。", Image: Image{AssetID: "plant-purple-basil", Alt: "晨光中的紫罗勒叶片"}},
		},
		Stories: []VolunteerStory{
			{ID: "story-mei", Name: "陈梅", Title: "把午休变成一小时的园艺课", Excerpt: "她记录下每次浇水后的土壤变化，也带新志愿者认识共享工具。", Image: Image{AssetID: "volunteer-chen-mei", Alt: "志愿者在高床边整理滴灌管"}},
			{ID: "story-bo", Name: "周博", Title: "和孩子们一起做昆虫旅馆", Excerpt: "旧竹筒和修枝材料有了新用途，观察卡也开始出现更多访客。", Image: Image{AssetID: "volunteer-zhou-bo", Alt: "亲子志愿者查看昆虫旅馆"}},
		},
		Activities: []Activity{
			{ID: "seedling-swap", Title: "秋播种苗交换会", Schedule: "8月24日 09:30-11:30", Location: "东门棚架", Capacity: 24, SignupURL: "/api/activities/registrations"},
			{ID: "night-insect-watch", Title: "夜间昆虫观察", Schedule: "8月30日 19:00-20:30", Location: "雨水花园", Capacity: 16, SignupURL: "/api/activities/registrations"},
		},
		Articles: []Article{
			{ID: "tomato-pruning", Title: "番茄侧枝修剪记录", Summary: "对比三周内不同修剪频率下的通风和挂果情况。", Category: "种植记录", Tags: []string{"番茄", "夏季", "高床"}, Image: Image{AssetID: "article-tomato-pruning", Alt: "番茄植株的修剪位置"}},
			{ID: "autumn-seeds", Title: "秋播种子发芽清单", Summary: "生菜、萝卜和芫荽的浸种与出芽记录。", Category: "种植记录", Tags: []string{"秋播", "育苗"}, Image: Image{AssetID: "article-autumn-seeds", Alt: "标记整齐的育苗盘"}},
			{ID: "ladybug-eggs", Title: "瓢虫卵观察第七天", Summary: "叶背卵块孵化后，蚜虫密度开始下降。", Category: "虫害观察", Tags: []string{"瓢虫", "蚜虫", "生态防治"}, Image: Image{AssetID: "article-ladybug-eggs", Alt: "叶片背面的瓢虫卵"}},
			{ID: "snail-route", Title: "雨后蜗牛活动路线", Summary: "用无伤害标记复查菜畦边缘的夜间移动路径。", Category: "虫害观察", Tags: []string{"蜗牛", "雨季", "夜间"}, Image: Image{AssetID: "article-snail-route", Alt: "湿润步道旁的蜗牛"}},
			{ID: "pruner-care", Title: "共享修枝剪归还检查", Summary: "借用后完成清洁、干燥和锁扣检查，延长工具寿命。", Category: "工具借用", Tags: []string{"修枝剪", "养护", "共享工具"}, Image: Image{AssetID: "article-pruner-care", Alt: "清洁后的共享修枝剪"}},
			{ID: "rain-barrel-map", Title: "邻里共绘雨水桶地图", Summary: "三栋楼的居民标记可收集雨水的位置和花园用水路线。", Category: "邻里共建", Tags: []string{"雨水", "地图", "节水"}, Image: Image{AssetID: "article-rain-barrel-map", Alt: "居民一起查看花园用水地图"}},
		},
		Categories: []string{"种植记录", "虫害观察", "工具借用", "邻里共建"},
	}
}
